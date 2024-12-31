package categories_contracts

import categories_entities "github.com/ladmakhi81/realtime-blogs/internal/categories/entities"

type CategoryRepositoryContract interface {
	CreateCategory(category *categories_entities.Category) error
	UpdateCategoryId()
	DeleteCategoryById()
	GetCategories()
	GetCategoryByTitle(title string) (*categories_entities.Category, error)
}
