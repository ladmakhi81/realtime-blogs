package auth_utils

import (
	"crypto/rand"
	"encoding/base64"
	"os"

	"github.com/golang-jwt/jwt/v5"
	auth_types "github.com/ladmakhi81/realtime-blogs/internal/auth/types"
	users_entities "github.com/ladmakhi81/realtime-blogs/internal/users/entities"
)

type JwtTokenGeneratorUtil struct{}

func NewJwtTokenGenerator() JwtTokenGeneratorUtil {
	return JwtTokenGeneratorUtil{}
}

func getTokenSecretKey() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func (JwtTokenGeneratorUtil) GenerateToken(user *users_entities.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, auth_types.NewLoggedInUserClaim(user.ID, user.Email))
	signedToken, err := token.SignedString(getTokenSecretKey())
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (JwtTokenGeneratorUtil) GenerateRefreshToken() (string, error) {
	refreshToken := make([]byte, 100)
	if _, err := rand.Read(refreshToken); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(refreshToken), nil
}
