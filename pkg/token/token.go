package token

import (
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/AkifhanIlgaz/hotel-booking-app/config"
	"github.com/AkifhanIlgaz/hotel-booking-app/internal/models"
	"github.com/AkifhanIlgaz/hotel-booking-app/migrations/queries"
	"github.com/AkifhanIlgaz/hotel-booking-app/pkg/errors"
	"github.com/AkifhanIlgaz/hotel-booking-app/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Manager struct {
	db                    *sql.DB
	accessTokenPrivateKey *rsa.PrivateKey
	accessTokenPublicKey  *rsa.PublicKey
	accessTokenExpiresIn  time.Duration
	refreshTokenExpiresIn time.Duration
}

type CustomClaims struct {
	jwt.RegisteredClaims
	Role string `json:"role"`
}

type ResetClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

func NewTokenManager(db *sql.DB, tokenConfig *config.TokenConfig) (*Manager, error) {
	tokenManager := Manager{
		db: db,
	}
	var err error

	tokenManager.accessTokenPrivateKey, err = loadPrivateKey(tokenConfig.PrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create token manager: %w", err)
	}

	tokenManager.accessTokenPublicKey, err = loadPublicKey(tokenConfig.PublicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create token manager: %w", err)
	}

	tokenManager.accessTokenExpiresIn = time.Duration(tokenConfig.AccessTokenExpiresIn) * time.Minute
	tokenManager.refreshTokenExpiresIn = time.Duration(tokenConfig.RefreshTokenExpiresIn) * time.Hour * 24

	return &tokenManager, nil
}

func (m *Manager) GenerateAccessToken(userId, role string) (string, error) {
	now := time.Now()

	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(m.accessTokenExpiresIn)),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   userId,
		},
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, err := token.SignedString(m.accessTokenPrivateKey)
	if err != nil {
		return "", fmt.Errorf("error signing access token: %w", err)
	}

	return signedToken, nil
}

func (m *Manager) GenerateResetToken(email string) (string, error) {
	now := time.Now()

	claims := ResetClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(2 * time.Minute)),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		Email: email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, err := token.SignedString(m.accessTokenPrivateKey)
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
			return m.accessTokenPublicKey, nil
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

func (m *Manager) ParseResetToken(resetToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		resetToken,
		&ResetClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// Verify that the signing algorithm is what we expect
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Return the public key for verification
			return m.accessTokenPublicKey, nil
		},
		// Additional validation options
		jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Name}),
		jwt.WithExpirationRequired(),
		jwt.WithIssuedAt(),
	)
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return "", fmt.Errorf("token expired: %w", err)
		case errors.Is(err, jwt.ErrTokenNotValidYet):
			return "", fmt.Errorf("token not valid yet: %w", err)
		default:
			return "", fmt.Errorf("invalid token: %w", err)
		}
	}

	// Verify token is valid and extract claims
	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	// Type assert the claims
	claims, ok := token.Claims.(*ResetClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	return claims.Email, nil
}

// TODO: Refactor => JWT refresh token, store hash of the token on Db
func (m *Manager) GenerateRefreshToken(uid uuid.UUID) (string, error) {
	expired, err := m.isRefreshTokenExpired(uid)
	if err != nil {
		return "", fmt.Errorf("error checking refresh token expiry: %w", err)
	}

	if expired {
		return "", errors.New("expired token")
	}

	token, err := utils.RandString(32)
	if err != nil {
		return "", fmt.Errorf("error generating refresh token: %w", err)
	}

	id := uuid.New()
	now := time.Now()
	expiry := now.Add(m.refreshTokenExpiresIn)

	var execErr error
	if errors.Is(err, sql.ErrNoRows) {
		_, execErr = m.db.Exec(queries.InsertRefreshToken,
			id,
			uid,
			utils.Hash(token),
			now,
			expiry,
		)
	} else {
		_, execErr = m.db.Exec(queries.UpdateRefreshToken, utils.Hash(token), now, expiry, uid)
	}

	if execErr != nil {
		return "", fmt.Errorf("db save error: %w", execErr)
	}

	return token, nil
}

func (m *Manager) ValidateRefreshToken(refreshToken string) (uuid.UUID, error) {
	var token models.RefreshToken

	err := m.db.QueryRow(queries.SelectRefreshToken, utils.Hash(refreshToken)).Scan(&token.Id, &token.UserId, &token.TokenHash, &token.ExpiresAt, &token.CreatedAt)
	if err != nil {
		// Todo: if sql.ErrNoRows return custom error
		return uuid.Nil, fmt.Errorf("check refresh token is expired: %w", err)
	}

	if time.Now().After(token.ExpiresAt) {
		return uuid.Nil, errors.ErrTokenExpired
	}

	return token.UserId, nil
}

func (m *Manager) DeleteRefreshToken(uid uuid.UUID) error {
	res, err := m.db.Exec(queries.DeleteRefreshToken, uid)
	if err != nil {
		return fmt.Errorf("error deleting refresh token: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errors.ErrNotFoundRefreshToken
	}

	return nil
}

func (m *Manager) isRefreshTokenExpired(uid uuid.UUID) (bool, error) {
	var expiry time.Time

	err := m.db.QueryRow(queries.ExpiryCheck, uid).Scan(&expiry)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("check refresh token is expired: %w", err)
	}

	return time.Now().After(expiry), nil
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
