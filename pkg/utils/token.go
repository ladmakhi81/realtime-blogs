package pkg_utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	pkg_types "github.com/ladmakhi81/realtime-blogs/pkg/types"
)

func getTokenSecretKey() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func GenerateToken(id uint, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, pkg_types.NewUserAuthClaim(id, email))
	signedToken, err := token.SignedString(getTokenSecretKey())
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func GenerateRefreshToken() (string, error) {
	refreshToken := make([]byte, 100)
	if _, err := rand.Read(refreshToken); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(refreshToken), nil
}

func VerifyAccessToken(token string) (*pkg_types.UserAuthClaim, error) {
	decodedToken, decodedTokenErr := jwt.ParseWithClaims(
		token,
		&pkg_types.UserAuthClaim{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid token method")
			}
			return getTokenSecretKey(), nil
		},
	)
	if decodedTokenErr != nil || !decodedToken.Valid {
		return nil, pkg_types.NewClientError(
			http.StatusUnauthorized,
			"Unauthorized",
		)
	}
	claims := decodedToken.Claims.(*pkg_types.UserAuthClaim)
	if isExpired := claims.ExpiresAt.Time.Before(time.Now()); isExpired {
		return nil, pkg_types.NewClientError(
			http.StatusUnauthorized,
			"Unauthorized",
		)
	}
	return claims, nil
}
