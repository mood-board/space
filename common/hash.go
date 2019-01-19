package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func ComputeHmac256(message string, secret string) (string, error) {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	_, err := h.Write([]byte(message))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
