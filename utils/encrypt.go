package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

// SrringToEncryption 文字列を暗号化する
func SrringToEncryption(s string) string {
	converted := sha256.Sum256([]byte(s))

	return hex.EncodeToString(converted[:])
}

// IntToEncryption 数値を暗号化する
func IntToEncryption(i int) string {
	s := strconv.Itoa(i)

	return SrringToEncryption(s)
}
