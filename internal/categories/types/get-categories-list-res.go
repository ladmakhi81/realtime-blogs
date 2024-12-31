package categories_types

import categories_entities "github.com/ladmakhi81/realtime-blogs/internal/categories/entities"

type GetCategoriesListResponse struct {
	Categories *[]categories_entities.Category `json:"categories"`
}