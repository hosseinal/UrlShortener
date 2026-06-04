package handler

import (
	"crypto/sha256"
	"encoding/hex"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generate256BitCode(input string) string {
	// sha256.Sum256 returns a fixed [32]byte array (32 bytes * 8 bits = 256 bits)
	hash := sha256.Sum256([]byte(input))

	// Convert the 32 bytes to a 64-character hex string
	return hex.EncodeToString(hash[:])
}
