package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hash password: %w", err)
	}

	return string(hashedBytes), nil
}

func HashRefreshToken(token string) string {
	hashedBytes := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hashedBytes[:])
}
