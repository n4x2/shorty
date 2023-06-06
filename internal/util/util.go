package util

import (
	"crypto/rand"
	"encoding/hex"
)

// ShortenURL generates a shortened URL based on the provided long URL.
// It generates a random hash and takes the first 8 characters of the hash as the shortened URL.
// If the baseURL environment variable is set, it is appended to the shortened URL.
func ShortenURL(longURL string) string {
	hash := generateRandomHash()

	// Take the first 8 characters of the hash as the shortened URL
	shortURL := hash[:8]

	return shortURL
}

func generateRandomHash() string {
	hash := make([]byte, 16)
	_, err := rand.Read(hash)
	if err != nil {
		panic(err) // Handle error appropriately in your application
	}
	return hex.EncodeToString(hash)
}
