package categories_contracts

import (
	categories_entities "github.com/ladmakhi81/realtime-blogs/internal/categories/entities"
	categories_types "github.com/ladmakhi81/realtime-blogs/internal/categories/types"
	users_entities "github.com/ladmakhi81/realtime-blogs/internal/users/entities"
)

type CategoryServiceContract interface {
	CreateCategory(
		reqBody categories_types.ModifyCategoryReqBody,
		creator *users_entities.User,
	) (*categories_entities.Category, error)

	GetCategories(page, limit uint) (*[]categories_entities.Category, error)
}
