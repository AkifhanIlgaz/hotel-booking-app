package token

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/AkifhanIlgaz/hotel-booking-app/config"
)

type Manager struct {
	AccessTokenPrivateKey *rsa.PrivateKey
	AccessTokenPublicKey  *rsa.PublicKey
	AccessTokenExpiresIn  time.Duration
	AccessTokenMaxAge     time.Duration

	RefreshTokenPrivateKey *rsa.PrivateKey
	RefreshTokenPublicKey  *rsa.PublicKey
	RefreshTokenExpiresIn  time.Duration
	RefreshTokenMaxAge     time.Duration
}

func NewTokenManager(tokenConfig *config.TokenConfig) (*Manager, error) {
	var tokenManager Manager
	var err error

	tokenManager.AccessTokenPrivateKey, err = loadPrivateKey(tokenConfig.AccessTokenPrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create token manager: %w", err)
	}

	tokenManager.AccessTokenPublicKey, err = loadPublicKey(tokenConfig.AccessTokenPublicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create token manager: %w", err)
	}

	tokenManager.AccessTokenExpiresIn = time.Duration(tokenConfig.AccessTokenExpiresIn) * time.Minute
	tokenManager.AccessTokenMaxAge = time.Duration(tokenConfig.AccessTokenMaxAge) * time.Minute

	tokenManager.RefreshTokenPrivateKey, err = loadPrivateKey(tokenConfig.RefreshTokenPrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create token manager: %w", err)
	}

	tokenManager.RefreshTokenPublicKey, err = loadPublicKey(tokenConfig.RefreshTokenPublicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create token manager: %w", err)
	}

	tokenManager.RefreshTokenExpiresIn = time.Duration(tokenConfig.RefreshTokenExpiresIn) * time.Hour * 24
	tokenManager.RefreshTokenMaxAge = time.Duration(tokenConfig.RefreshTokenMaxAge) * time.Hour * 24

	return &tokenManager, nil
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
