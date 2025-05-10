package handlers

import (
	"time"
)

// /timestamp: Retorna la hora actual en formato ISO-8601.
func Timestamp() string {
	return time.Now().Format(time.RFC3339)
}