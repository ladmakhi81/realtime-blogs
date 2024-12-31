package auth_contracts

import users_entities "github.com/ladmakhi81/realtime-blogs/internal/users/entities"

type JwtTokenGeneratorContract interface {
	GenerateToken(user *users_entities.User) (string, error)
	GenerateRefreshToken() (string, error)
}
