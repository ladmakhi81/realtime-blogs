package categories_contracts

import categories_entities "github.com/ladmakhi81/realtime-blogs/internal/categories/entities"

type CategoryRepositoryContract interface {
	CreateCategory(category *categories_entities.Category) error
	UpdateCategoryId()
	DeleteCategoryById(id uint) error
	GetCategories(page, limit uint) (*[]categories_entities.Category, error)
	GetCategoryByTitle(title string) (*categories_entities.Category, error)
	GetCategoryById(id uint) (*categories_entities.Category, error)
}
