package auth_entities

import (
	users_entities "github.com/ladmakhi81/realtime-blogs/internal/users/entities"
	pkg_entities "github.com/ladmakhi81/realtime-blogs/pkg/entities"
)

type Token struct {
	AccessToken  string               `json:"accessToken"`
	RefreshToken string               `json:"refreshToken"`
	User         *users_entities.User `json:"user"`
	pkg_entities.BaseEntity
}
