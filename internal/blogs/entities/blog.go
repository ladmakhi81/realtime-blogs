package blogs_entities

import (
	categories_entities "github.com/ladmakhi81/realtime-blogs/internal/categories/entities"
	users_entities "github.com/ladmakhi81/realtime-blogs/internal/users/entities"
	pkg_entities "github.com/ladmakhi81/realtime-blogs/pkg/entities"
)

type Blog struct {
	Title     string                       `json:"title"`
	Content   string                       `json:"content"`
	CreatedBy users_entities.User          `json:"createdBy"`
	Category  categories_entities.Category `json:"category"`
	Tags      []string                     `json:"tags"`
	pkg_entities.BaseEntity
}

func NewBlog(
	title,
	content string,
	createdBy *users_entities.User,
	category *categories_entities.Category,
	tags []string,
) *Blog {
	return &Blog{
		Title:     title,
		Content:   content,
		CreatedBy: *createdBy,
		Category:  *category,
		Tags:      tags,
	}
}
