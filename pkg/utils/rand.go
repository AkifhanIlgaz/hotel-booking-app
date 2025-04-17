package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// RandString generates a random string of a specified length.
//
// It uses crypto/rand to generate cryptographically secure random bytes
// and then encodes them using base64 URL encoding.
//
// Parameters:
//   - n: The number of bytes to generate. The resulting string length will be longer due to base64 encoding.
//
// Returns:
//   - string: The generated random string.
//   - error: An error if the random bytes generation fails.
func RandString(n int) (string, error) {
	bytes := make([]byte, n) // 256-bit (32 * 8)

	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("generate random string: %w", err)
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}
