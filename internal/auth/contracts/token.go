package auth_contracts

import auth_entities "github.com/ladmakhi81/realtime-blogs/internal/auth/entities"

type TokenRepositoryContract interface {
	CreateToken(token *auth_entities.Token) error
	DeleteTokensByUserId()
	GetTokenByUserId()
}
