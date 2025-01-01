package pkg_types

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserAuthClaim struct {
	ID    uint
	Email string
	jwt.RegisteredClaims
}

func NewUserAuthClaim(id uint, email string) UserAuthClaim {
	return UserAuthClaim{
		ID:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "custom-token",
		},
	}
}
