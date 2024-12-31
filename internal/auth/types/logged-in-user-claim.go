package auth_types

import "github.com/golang-jwt/jwt/v5"

type LoggedInUserClaim struct {
	ID    uint
	Email string
	jwt.RegisteredClaims
}

func NewLoggedInUserClaim(id uint, email string) LoggedInUserClaim {
	return LoggedInUserClaim{ID: id, Email: email}
}
