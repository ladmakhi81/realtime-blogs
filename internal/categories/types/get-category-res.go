package categories_types

import categories_entities "github.com/ladmakhi81/realtime-blogs/internal/categories/entities"

type GetCategoryResponse struct {
	Category categories_entities.Category `json:"category"`
}
