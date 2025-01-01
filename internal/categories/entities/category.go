package categories_entities

import (
	users_entities "github.com/ladmakhi81/realtime-blogs/internal/users/entities"
	pkg_entities "github.com/ladmakhi81/realtime-blogs/pkg/entities"
)

type Category struct {
	Title     string               `json:"title,omitempty"`
	CreatedBy *users_entities.User `json:"createdBy,omitempty"`
	pkg_entities.BaseEntity
}

func NewCategory(title string, createdBy *users_entities.User) *Category {
	return &Category{Title: title, CreatedBy: createdBy}
}
