package pkg_utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateCode(len uint) (string, error) {
	refreshToken := make([]byte, len)
	if _, err := rand.Read(refreshToken); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(refreshToken), nil
}
