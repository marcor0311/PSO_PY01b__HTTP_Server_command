package handlers

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"httpserver/internal/utils"
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

// /countwords: Counts lowercase words and ignores punctuation/symbols
func CountWords(text string) map[string]int {
	freq := make(map[string]int)
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		raw := scanner.Text()
		clean := utils.FilterLettersOnly(raw)
		if clean != "" {
			freq[strings.ToLower(clean)]++
		}
	}
	return freq
}
