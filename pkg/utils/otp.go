package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// GenerateNumericOTP generates a random numeric OTP of specified length
func GenerateNumericOTP(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("invalid length: %d", length)
	}

	otp := make([]byte, length)
	for i := range length {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", fmt.Errorf("generate random number: %w", err)
		}
		otp[i] = byte(num.Int64() + '0')
	}

	return string(otp), nil
}
