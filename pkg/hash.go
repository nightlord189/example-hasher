package pkg

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func GetSHA256Hash(plainText string) string {
	h := sha256.New()
	h.Write([]byte(plainText))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GetSHA512Hash(plainText string) string {
	h := sha512.New()
	h.Write([]byte(plainText))
	return fmt.Sprintf("%x", h.Sum(nil))
}
