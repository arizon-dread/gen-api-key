package services

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(original string) string {
	bytes := []byte(original)
	res := sha256.Sum256(bytes)

	return hex.EncodeToString(res[:])
}
