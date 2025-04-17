package token

import (
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/AkifhanIlgaz/hotel-booking-app/config"
	"github.com/AkifhanIlgaz/hotel-booking-app/migrations/queries"
	"github.com/AkifhanIlgaz/hotel-booking-app/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Manager struct {
	db                    *sql.DB
	AccessTokenPrivateKey *rsa.PrivateKey
	AccessTokenPublicKey  *rsa.PublicKey
	AccessTokenExpiresIn  time.Duration
	RefreshTokenExpiresIn time.Duration
}

type CustomClaims struct {
	jwt.RegisteredClaims
	Role string `json:"role"`
}

func NewTokenManager(db *sql.DB, tokenConfig *config.TokenConfig) (*Manager, error) {
	tokenManager := Manager{
		db: db,
	}
	var err error

	tokenManager.AccessTokenPrivateKey, err = loadPrivateKey(tokenConfig.PrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create token manager: %w", err)
	}

	tokenManager.AccessTokenPublicKey, err = loadPublicKey(tokenConfig.PublicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create token manager: %w", err)
	}

	tokenManager.AccessTokenExpiresIn = time.Duration(tokenConfig.AccessTokenExpiresIn) * time.Minute
	tokenManager.RefreshTokenExpiresIn = time.Duration(tokenConfig.RefreshTokenExpiresIn) * time.Hour * 24

	return &tokenManager, nil
}

func (m *Manager) GenerateAccessToken(userId, role string) (string, error) {
	now := time.Now()

	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(m.AccessTokenExpiresIn)),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   userId,
		},
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, err := token.SignedString(m.AccessTokenPrivateKey)
	if err != nil {
		return "", fmt.Errorf("error signing access token: %w", err)
	}

	return signedToken, nil
}

func (m *Manager) ParseAccessToken(accessToken string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// Verify that the signing algorithm is what we expect
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Return the public key for verification
			return m.AccessTokenPublicKey, nil
		},
		// Additional validation options
		jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Name}),
		jwt.WithExpirationRequired(),
		jwt.WithIssuedAt(),
	)
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, fmt.Errorf("token expired: %w", err)
		case errors.Is(err, jwt.ErrTokenNotValidYet):
			return nil, fmt.Errorf("token not valid yet: %w", err)
		default:
			return nil, fmt.Errorf("invalid token: %w", err)
		}
	}

	// Verify token is valid and extract claims
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Type assert the claims
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func (m *Manager) GenerateRefreshToken(uid uuid.UUID) (string, error) {
	token, err := utils.RandString(32)
	if err != nil {
		return "", fmt.Errorf("error generating refresh token: %w", err)
	}

	id := uuid.New()
	now := time.Now()
	expiry := now.Add(m.RefreshTokenExpiresIn * 24 * time.Hour)

	if _, err := m.db.Exec(queries.InsertRefreshToken,
		id,
		uid,
		utils.HashRefreshToken(token),
		now,
		expiry,
	); err != nil {
		return "", fmt.Errorf("generate refresh token: %w", err)
	}

	return token, nil
}

func loadPrivateKey(path string) (*rsa.PrivateKey, error) {
	privateKeyBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read private key file: %w", err)
	}

	block, _ := pem.Decode(privateKeyBytes)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing RSA private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return privateKey, nil
}

func loadPublicKey(path string) (*rsa.PublicKey, error) {
	publicKeyBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read public key file: %w", err)
	}

	// Decode PEM block
	block, _ := pem.Decode(publicKeyBytes)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	// Parse the public key
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	// Assert that the key is an RSA public key
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("key is not an RSA public key")
	}

	return rsaPublicKey, nil
}
