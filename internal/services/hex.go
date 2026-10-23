// Package services contains internal service funcs for the cli app
package services

import (
	"crypto/rand"
	"encoding/hex"
)

func RandomHex() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
