package utils

import (
	"httpserver/internal/constants"
	"os"
	"strings"
)

func GetEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return def
}

func IsParallel(path string) bool {
    base := strings.SplitN(path, "?", 2)[0] 
    for _, r := range constants.ParallelRoutes {
        if base == r {
            return true
        }
    }
    return false
}