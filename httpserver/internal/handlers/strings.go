package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// /reverse?text=abc: Returns the input string reversed.
func ReverseString(input string) (reversed string) {
	for _, char := range input {
		reversed = string(char) + reversed
	}
	return
}

// /toupper?text=abc: Converts the text to uppercase.
func ToUpper(text string) string {
	return strings.ToUpper(text)
}

// /hash?text=abc: Returns the SHA-256 hash of the text
func HashSHA256(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}