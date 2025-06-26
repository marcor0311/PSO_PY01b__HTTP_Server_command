package utils

import (
	"httpserver/internal/constants"
	"os"
	"strings"
	"unicode"
)

/**
 * Retrieves an environment variable or returns a default value if it isn’t set.
 * @param {string} key - Name of the environment variable.
 * @param {string} def - Fallback value when the variable is absent or empty.
 */
func GetEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return def
}

/**
 * Determines whether a given path is marked for parallel processing.
 * @param {string} path - Request path.
 */
func IsParallel(path string) bool {
	base := strings.SplitN(path, "?", 2)[0]
	for _, r := range constants.ParallelRoutes {
		if base == r {
			return true
		}
	}
	return false
}

// filterLettersOnly removes all non-letter characters from a word.
func FilterLettersOnly(word string) string {
	var b strings.Builder
	for _, r := range word {
		if unicode.IsLetter(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}
