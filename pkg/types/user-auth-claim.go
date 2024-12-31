package pkg_types

import "github.com/golang-jwt/jwt/v5"

type UserAuthClaim struct {
	ID    uint
	Email string
	jwt.RegisteredClaims
}

func NewUserAuthClaim(id uint, email string) UserAuthClaim {
	return UserAuthClaim{ID: id, Email: email}
}
