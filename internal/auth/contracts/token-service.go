package auth_contracts

import (
	auth_entities "github.com/ladmakhi81/realtime-blogs/internal/auth/entities"
	users_entities "github.com/ladmakhi81/realtime-blogs/internal/users/entities"
)

type TokenServiceContractor interface {
	CreateToken(user *users_entities.User) (*auth_entities.Token, error)
}
