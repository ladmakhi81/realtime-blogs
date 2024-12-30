package categories_services

import categories_contracts "github.com/ladmakhi81/realtime-blogs/internal/categories/contracts"

type CategoryService struct {
	categoryRepo categories_contracts.CategoryRepositoryContract
}

func (categoryService CategoryService) CreateCategory() {}

func (categoryService CategoryService) DeleteCategoryById() {}

func (categoryService CategoryService) UpdateCategoryById() {}

func (categoryService CategoryService) GetCategories() {}
