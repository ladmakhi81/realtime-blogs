package pkg_utils

import (
	users_entities "github.com/ladmakhi81/realtime-blogs/internal/users/entities"
)

func GenerateToken(user *users_entities.User) (string, error) {
	return "access token D)", nil
}

func GenerateRefreshToken() (string, error) {
	return "refresh token D)", nil
	// refreshToken := make([]byte, len)
	// if _, err := rand.Read(refreshToken); err != nil {
	// 	return "", err
	// }
	// return base64.URLEncoding.EncodeToString(refreshToken[:]), nil
}
