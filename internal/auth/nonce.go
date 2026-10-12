package auth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"sync"
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
)

const (
	nonceRandomBytes = 16
	nonceExpiryBytes = 8
	nonceMACBytes    = 16
	nonceBytes       = nonceRandomBytes + nonceExpiryBytes + nonceMACBytes

	usedNonceCapacity = 100_000
)

func (s *Service) Nonce() (string, time.Time, error) {
	nonce := make([]byte, nonceBytes)

	// [random][expiry][mac]
	if _, err := rand.Read(nonce[:nonceRandomBytes]); err != nil {
		return "", time.Time{}, fmt.Errorf("%w: %w", ErrNonceGeneration, err)
	}

	expires := s.now().Add(NonceTTL).Truncate(time.Second)
	binary.BigEndian.PutUint64(nonce[nonceRandomBytes:], uint64(expires.Unix()))
	copy(nonce[nonceRandomBytes+nonceExpiryBytes:], s.nonceMAC(nonce[:nonceRandomBytes+nonceExpiryBytes]))

	return hex.EncodeToString(nonce), expires, nil
}

func (s *Service) checkNonce(nonce string) (string, error) {
	raw, err := hex.DecodeString(nonce)
	if err != nil || len(raw) != nonceBytes {
		return "", ErrInvalidNonce
	}

	payload, mac := raw[:nonceRandomBytes+nonceExpiryBytes], raw[nonceRandomBytes+nonceExpiryBytes:]
	if !hmac.Equal(mac, s.nonceMAC(payload)) {
		return "", ErrInvalidNonce
	}

	expires := time.Unix(int64(binary.BigEndian.Uint64(raw[nonceRandomBytes:])), 0)
	if !s.now().Before(expires) {
		return "", ErrNonceExpired
	}
	return hex.EncodeToString(raw), nil
}

func (s *Service) nonceMAC(payload []byte) []byte {
	mac := hmac.New(sha256.New, s.nonceKey)
	mac.Write(payload)
	return mac.Sum(nil)[:nonceMACBytes]
}

type usedNonces struct {
	mu    sync.Mutex
	cache *expirable.LRU[string, struct{}]
}

func newUsedNonces() *usedNonces {
	return &usedNonces{
		cache: expirable.NewLRU[string, struct{}](usedNonceCapacity, nil, NonceTTL),
	}
}

func (u *usedNonces) markUsed(nonce string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()

	if u.cache.Contains(nonce) {
		return false
	}
	u.cache.Add(nonce, struct{}{})
	return true
}
