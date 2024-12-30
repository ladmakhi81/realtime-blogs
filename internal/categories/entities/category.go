package categories_entities

import (
	users_entities "github.com/ladmakhi81/realtime-blogs/internal/users/entities"
	pkg_entities "github.com/ladmakhi81/realtime-blogs/pkg/entities"
)

type Category struct {
	Title     string               `json:"title"`
	CreatedBy *users_entities.User `json:"createdBy"`
	pkg_entities.BaseEntity
}
