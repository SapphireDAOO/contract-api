package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	NonceTTL        = 5 * time.Minute
	SessionTTL      = 12 * time.Hour
	MinSecretLength = 32

	issuer = "sapphire-contract-api"
)

type Service struct {
	domains  map[string]bool
	chainID  int64
	nonceKey []byte
	tokenKey []byte
	used     *usedNonces
	now      func() time.Time
}

func New(secret string, domains []string, chainID int64) (*Service, error) {
	if len(secret) < MinSecretLength {
		return nil, ErrSecretTooShort
	}

	allowed := make(map[string]bool, len(domains))
	for _, domain := range domains {
		if domain = strings.TrimSpace(domain); domain != "" {
			allowed[domain] = true
		}
	}
	if len(allowed) == 0 {
		return nil, ErrNoDomains
	}

	return &Service{
		domains:  allowed,
		chainID:  chainID,
		nonceKey: deriveKey(secret, "sapphire/siwe-nonce/v1"),
		tokenKey: deriveKey(secret, "sapphire/session/v1"),
		used:     newUsedNonces(),
		now:      time.Now,
	}, nil
}

func deriveKey(secret, purpose string) []byte {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(purpose))
	return mac.Sum(nil)
}

type addressKey struct{}

func WithAddress(ctx context.Context, address common.Address) context.Context {
	return context.WithValue(ctx, addressKey{}, address)
}

func AddressFrom(ctx context.Context) (common.Address, bool) {
	address, ok := ctx.Value(addressKey{}).(common.Address)
	return address, ok
}
