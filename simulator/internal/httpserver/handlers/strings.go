package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strings"
)

// /reverse?text=abc: Invierte el texto recibido.
func HandleReverse(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "Missing 'text' parameter", http.StatusBadRequest)
		return
	}
	reversed := reverseString(text)
	w.Write([]byte(reversed))
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// /toupper?text=abc: Convierte el texto a mayusculas.
func HandleToUpper(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "Missing 'text' parameter", http.StatusBadRequest)
		return
	}
	w.Write([]byte(strings.ToUpper(text)))
}

// /hash?text=abc: Devuelve el SHA-256 del texto
func HandleHash(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "Missing 'text' parameter", http.StatusBadRequest)
		return
	}
	hash := sha256.Sum256([]byte(text))
	w.Write([]byte(hex.EncodeToString(hash[:])))
}