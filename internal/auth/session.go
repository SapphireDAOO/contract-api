package auth

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spruceid/siwe-go"
)

type Session struct {
	Address   common.Address
	Token     string
	ExpiresAt time.Time
}

func (s *Service) SignIn(message, signature string) (*Session, error) {
	msg, err := siwe.ParseMessage(message)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidMessage, err)
	}

	if !s.domains[msg.GetDomain()] {
		return nil, ErrWrongDomain
	}

	if int64(msg.GetChainID()) != s.chainID {
		return nil, ErrWrongChain
	}

	if _, err := msg.ValidAt(s.now()); err != nil {
		return nil, ErrMessageExpired
	}

	nonce, err := s.checkNonce(msg.GetNonce())
	if err != nil {
		return nil, err
	}

	signer, err := recoverSigner(message, signature)
	if err != nil || signer != msg.GetAddress() {
		return nil, ErrInvalidSignature
	}

	if !s.used.markUsed(nonce) {
		return nil, ErrNonceUsed
	}
	return s.issueSession(signer)
}

func recoverSigner(message, signature string) (common.Address, error) {
	sig, err := hexutil.Decode(signature)
	if err != nil || len(sig) != crypto.SignatureLength {
		return common.Address{}, ErrInvalidSignature
	}
	if sig[crypto.RecoveryIDOffset] >= 27 {
		sig[crypto.RecoveryIDOffset] -= 27
	}

	key, err := crypto.SigToPub(accounts.TextHash([]byte(message)), sig)
	if err != nil {
		return common.Address{}, ErrInvalidSignature
	}
	return crypto.PubkeyToAddress(*key), nil
}

type sessionClaims struct {
	ChainID int64 `json:"chainId"`
	jwt.RegisteredClaims
}

func (s *Service) issueSession(address common.Address) (*Session, error) {
	now := s.now()
	expires := now.Add(SessionTTL)

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, sessionClaims{
		ChainID: s.chainID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			Subject:   address.Hex(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expires),
		},
	}).SignedString(s.tokenKey)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrSessionSigning, err)
	}

	return &Session{Address: address, Token: token, ExpiresAt: expires}, nil
}

func (s *Service) SessionAddress(token string) (common.Address, error) {
	var claims sessionClaims
	_, err := jwt.ParseWithClaims(token, &claims,
		func(*jwt.Token) (any, error) { return s.tokenKey, nil },
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithExpirationRequired(),
		jwt.WithIssuer(issuer),
		jwt.WithTimeFunc(s.now),
	)
	if err != nil || claims.ChainID != s.chainID || !common.IsHexAddress(claims.Subject) {
		return common.Address{}, ErrInvalidSession
	}
	return common.HexToAddress(claims.Subject), nil
}
