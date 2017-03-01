package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"strconv"

	"github.com/astaxie/beego"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

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

// Encrypter 暗号化する
func Encrypter(s []byte) (string, error) {
	c, err := aes.NewCipher([]byte(beego.AppConfig.String("encryptionKey")))
	if err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(s))
	cfb.XORKeyStream(ciphertext, s)

	return string(ciphertext), nil
}

// Decrypter 復号化する
func Decrypter(s []byte) (string, error) {
	c, err := aes.NewCipher([]byte(beego.AppConfig.String("encryptionKey")))
	if err != nil {
		return "", err
	}

	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(s))
	cfbdec.XORKeyStream(plaintextCopy, s)

	return string(plaintextCopy), nil
}
