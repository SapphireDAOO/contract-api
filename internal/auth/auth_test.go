package auth

import (
	"crypto/ecdsa"
	"errors"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spruceid/siwe-go"
)

const (
	testSecret  = "0123456789abcdef0123456789abcdef0123456789abcdef"
	testDomain  = "app.example.com"
	testChainID = 84532
)

type testClock struct{ now time.Time }

func (c *testClock) Now() time.Time { return c.now }

func newTestService(t *testing.T) (*Service, *testClock) {
	t.Helper()
	s, err := New(testSecret, []string{testDomain}, testChainID)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	clock := &testClock{now: time.Date(2026, 10, 12, 12, 0, 0, 0, time.UTC)}
	s.now = clock.Now
	return s, clock
}

func newKey(t *testing.T) *ecdsa.PrivateKey {
	t.Helper()
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("generating key: %v", err)
	}
	return key
}

type messageOptions struct {
	domain  string
	chainID int
	expires *time.Time
}

func signedMessage(t *testing.T, key *ecdsa.PrivateKey, nonce string, opts messageOptions) (string, string) {
	t.Helper()
	if opts.domain == "" {
		opts.domain = testDomain
	}
	if opts.chainID == 0 {
		opts.chainID = testChainID
	}

	options := map[string]any{
		"chainId":   opts.chainID,
		"statement": "Sign in to write and read invoice notes.",
		"issuedAt":  time.Date(2026, 10, 12, 11, 59, 0, 0, time.UTC),
	}
	if opts.expires != nil {
		options["expirationTime"] = *opts.expires
	}

	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	msg, err := siwe.InitMessage(opts.domain, address, "https://"+opts.domain, nonce, options)
	if err != nil {
		t.Fatalf("building message: %v", err)
	}
	text := msg.String()

	sig, err := crypto.Sign(accounts.TextHash([]byte(text)), key)
	if err != nil {
		t.Fatalf("signing: %v", err)
	}
	sig[crypto.RecoveryIDOffset] += 27
	return text, hexutil.Encode(sig)
}

func issueNonce(t *testing.T, s *Service) string {
	t.Helper()
	nonce, _, err := s.Nonce()
	if err != nil {
		t.Fatalf("Nonce: %v", err)
	}
	return nonce
}

func TestNewValidatesItsInputs(t *testing.T) {
	if _, err := New("short", []string{testDomain}, testChainID); !errors.Is(err, ErrSecretTooShort) {
		t.Errorf("short secret: err = %v, want ErrSecretTooShort", err)
	}
	if _, err := New(testSecret, []string{"", "  "}, testChainID); !errors.Is(err, ErrNoDomains) {
		t.Errorf("no usable domain: err = %v, want ErrNoDomains", err)
	}
}

func TestNonce(t *testing.T) {
	s, clock := newTestService(t)

	nonce, expires, err := s.Nonce()
	if err != nil {
		t.Fatalf("Nonce: %v", err)
	}
	if want := clock.now.Add(NonceTTL); !expires.Equal(want) {
		t.Errorf("expires = %v, want %v", expires, want)
	}
	if strings.Trim(nonce, "0123456789abcdef") != "" || len(nonce) < 8 {
		t.Errorf("nonce %q is not EIP-4361 alphanumeric", nonce)
	}
	if _, err := s.checkNonce(nonce); err != nil {
		t.Errorf("a fresh nonce was rejected: %v", err)
	}

	tampered := []byte(nonce)
	tampered[0] ^= 1
	if _, err := s.checkNonce(string(tampered)); !errors.Is(err, ErrInvalidNonce) {
		t.Errorf("tampered nonce: err = %v, want ErrInvalidNonce", err)
	}

	other, err := New(strings.Repeat("z", MinSecretLength), []string{testDomain}, testChainID)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if _, err := other.checkNonce(nonce); !errors.Is(err, ErrInvalidNonce) {
		t.Errorf("nonce from another secret: err = %v, want ErrInvalidNonce", err)
	}

	clock.now = clock.now.Add(NonceTTL)
	if _, err := s.checkNonce(nonce); !errors.Is(err, ErrNonceExpired) {
		t.Errorf("expired nonce: err = %v, want ErrNonceExpired", err)
	}
}

func TestSignIn(t *testing.T) {
	s, clock := newTestService(t)
	key := newKey(t)
	message, signature := signedMessage(t, key, issueNonce(t, s), messageOptions{})

	session, err := s.SignIn(message, signature)
	if err != nil {
		t.Fatalf("SignIn: %v", err)
	}

	want := crypto.PubkeyToAddress(key.PublicKey)
	if session.Address != want {
		t.Errorf("address = %s, want %s", session.Address, want)
	}
	if !session.ExpiresAt.Equal(clock.now.Add(SessionTTL)) {
		t.Errorf("expiresAt = %v, want %v", session.ExpiresAt, clock.now.Add(SessionTTL))
	}

	got, err := s.SessionAddress(session.Token)
	if err != nil {
		t.Fatalf("SessionAddress: %v", err)
	}
	if got != want {
		t.Errorf("session address = %s, want %s", got, want)
	}
}

func TestSignInRejects(t *testing.T) {
	expired := time.Date(2026, 10, 12, 11, 59, 30, 0, time.UTC)

	tests := []struct {
		name  string
		build func(t *testing.T, s *Service, key *ecdsa.PrivateKey) (string, string)
		want  error
	}{
		{"another domain", func(t *testing.T, s *Service, key *ecdsa.PrivateKey) (string, string) {
			return signedMessage(t, key, issueNonce(t, s), messageOptions{domain: "evil.example.com"})
		}, ErrWrongDomain},
		{"another chain", func(t *testing.T, s *Service, key *ecdsa.PrivateKey) (string, string) {
			return signedMessage(t, key, issueNonce(t, s), messageOptions{chainID: 8453})
		}, ErrWrongChain},
		{"expired message", func(t *testing.T, s *Service, key *ecdsa.PrivateKey) (string, string) {
			return signedMessage(t, key, issueNonce(t, s), messageOptions{expires: &expired})
		}, ErrMessageExpired},
		{"nonce not issued here", func(t *testing.T, s *Service, key *ecdsa.PrivateKey) (string, string) {
			return signedMessage(t, key, "abcdef0123456789", messageOptions{})
		}, ErrInvalidNonce},
		{"signed by another key", func(t *testing.T, s *Service, key *ecdsa.PrivateKey) (string, string) {
			message, _ := signedMessage(t, key, issueNonce(t, s), messageOptions{})
			sig, err := crypto.Sign(accounts.TextHash([]byte(message)), newKey(t))
			if err != nil {
				t.Fatalf("signing: %v", err)
			}
			return message, hexutil.Encode(sig)
		}, ErrInvalidSignature},
		{"signature too short", func(t *testing.T, s *Service, key *ecdsa.PrivateKey) (string, string) {
			message, _ := signedMessage(t, key, issueNonce(t, s), messageOptions{})
			return message, "0x1234"
		}, ErrInvalidSignature},
		{"signature is not hex", func(t *testing.T, s *Service, key *ecdsa.PrivateKey) (string, string) {
			message, _ := signedMessage(t, key, issueNonce(t, s), messageOptions{})
			return message, "not-a-signature"
		}, ErrInvalidSignature},
		{"message changed after signing", func(t *testing.T, s *Service, key *ecdsa.PrivateKey) (string, string) {
			message, signature := signedMessage(t, key, issueNonce(t, s), messageOptions{})
			return strings.Replace(message, "invoice notes", "invoice notez", 1), signature
		}, ErrInvalidSignature},
		{"not a SIWE message", func(t *testing.T, s *Service, key *ecdsa.PrivateKey) (string, string) {
			return "hello", "0x00"
		}, ErrInvalidMessage},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := newTestService(t)
			message, signature := tt.build(t, s, newKey(t))

			if _, err := s.SignIn(message, signature); !errors.Is(err, tt.want) {
				t.Errorf("err = %v, want %v", err, tt.want)
			}
		})
	}
}

func TestSignInNonceIsSingleUse(t *testing.T) {
	s, _ := newTestService(t)
	message, signature := signedMessage(t, newKey(t), issueNonce(t, s), messageOptions{})

	if _, err := s.SignIn(message, signature); err != nil {
		t.Fatalf("first SignIn: %v", err)
	}
	if _, err := s.SignIn(message, signature); !errors.Is(err, ErrNonceUsed) {
		t.Errorf("replayed SignIn: err = %v, want ErrNonceUsed", err)
	}
}

func TestSignInSpendsTheNonceOnlyOnSuccess(t *testing.T) {
	s, _ := newTestService(t)
	key := newKey(t)
	message, signature := signedMessage(t, key, issueNonce(t, s), messageOptions{})

	forged, err := crypto.Sign(accounts.TextHash([]byte(message)), newKey(t))
	if err != nil {
		t.Fatalf("signing: %v", err)
	}
	if _, err := s.SignIn(message, hexutil.Encode(forged)); !errors.Is(err, ErrInvalidSignature) {
		t.Fatalf("forged SignIn: err = %v, want ErrInvalidSignature", err)
	}
	if _, err := s.SignIn(message, signature); err != nil {
		t.Errorf("a failed attempt spent the nonce: %v", err)
	}
}

func TestSessionAddressRejects(t *testing.T) {
	s, clock := newTestService(t)
	session, err := s.issueSession(common.HexToAddress("0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed"))
	if err != nil {
		t.Fatalf("issueSession: %v", err)
	}

	otherSecret, err := New(strings.Repeat("z", MinSecretLength), []string{testDomain}, testChainID)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	otherChain, err := New(testSecret, []string{testDomain}, 8453)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	unsigned, err := jwt.NewWithClaims(jwt.SigningMethodNone, sessionClaims{
		ChainID: testChainID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			Subject:   "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed",
			ExpiresAt: jwt.NewNumericDate(clock.now.Add(time.Hour)),
		},
	}).SignedString(jwt.UnsafeAllowNoneSignatureType)
	if err != nil {
		t.Fatalf("building an unsigned token: %v", err)
	}

	for name, check := range map[string]func() error{
		"another secret": func() error { _, err := otherSecret.SessionAddress(session.Token); return err },
		"another chain":  func() error { _, err := otherChain.SessionAddress(session.Token); return err },
		"unsigned token": func() error { _, err := s.SessionAddress(unsigned); return err },
		"garbage":        func() error { _, err := s.SessionAddress("not.a.token"); return err },
		"expired": func() error {
			clock.now = clock.now.Add(SessionTTL)
			defer func() { clock.now = clock.now.Add(-SessionTTL) }()
			_, err := s.SessionAddress(session.Token)
			return err
		},
	} {
		if err := check(); !errors.Is(err, ErrInvalidSession) {
			t.Errorf("%s: err = %v, want ErrInvalidSession", name, err)
		}
	}
}

func TestMarkUsedAllowsOneWinner(t *testing.T) {
	used := newUsedNonces()
	var wins atomic.Int32
	var wg sync.WaitGroup

	for range 50 {
		wg.Go(func() {
			if used.markUsed("same-nonce") {
				wins.Add(1)
			}
		})
	}
	wg.Wait()

	if got := wins.Load(); got != 1 {
		t.Errorf("%d callers spent the same nonce, want 1", got)
	}
}
