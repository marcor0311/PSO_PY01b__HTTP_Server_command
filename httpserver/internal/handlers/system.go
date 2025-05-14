package handlers

import (
	"time"
)

// /timestamp: Returns the current time.
func Timestamp() string {
	return time.Now().Format(time.RFC3339)
}