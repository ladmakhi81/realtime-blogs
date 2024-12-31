package categories_services

import (
	"net/http"

	categories_contracts "github.com/ladmakhi81/realtime-blogs/internal/categories/contracts"
	categories_entities "github.com/ladmakhi81/realtime-blogs/internal/categories/entities"
	categories_types "github.com/ladmakhi81/realtime-blogs/internal/categories/types"
	users_entities "github.com/ladmakhi81/realtime-blogs/internal/users/entities"
	pkg_types "github.com/ladmakhi81/realtime-blogs/pkg/types"
)

type CategoryService struct {
	CategoryRepo categories_contracts.CategoryRepositoryContract
}

func NewCategoryService(
	categoryRepo categories_contracts.CategoryRepositoryContract,
) CategoryService {
	return CategoryService{CategoryRepo: categoryRepo}
}

func (categoryService CategoryService) CreateCategory(
	reqBody categories_types.ModifyCategoryReqBody,
	creator *users_entities.User,
) (*categories_entities.Category, error) {
	duplicatedCategory, duplicatedCategoryErr := categoryService.CategoryRepo.GetCategoryByTitle(reqBody.Title)
	if duplicatedCategoryErr != nil {
		return nil, pkg_types.NewServerError(
			"error in finding category by title",
			"CategoryService.CreateCategory.GetCategoryByTitle",
			duplicatedCategoryErr.Error(),
		)
	}
	if duplicatedCategory != nil {
		return nil, pkg_types.NewClientError(
			http.StatusConflict,
			"category already exist by this title",
		)
	}
	category := categories_entities.NewCategory(reqBody.Title, creator)
	if createErr := categoryService.CategoryRepo.CreateCategory(category); createErr != nil {
		return nil, pkg_types.NewServerError(
			"error in creating category",
			"CategoryService.CreateCategory.CreateCategory",
			createErr.Error(),
		)
	}
	return category, nil
}

func (categoryService CategoryService) GetCategories(page, limit uint) (*[]categories_entities.Category, error) {
	categories, categoriesErr := categoryService.CategoryRepo.GetCategories(page, limit)
	if categoriesErr != nil {
		return nil, pkg_types.NewServerError(
			"unable to fetch categories from database",
			"CategoryService.GetCategories",
			categoriesErr.Error(),
		)
	}
	return categories, nil
}

func (categoryService CategoryService) GetCategoryById(id uint) (*categories_entities.Category, error) {
	category, categoryErr := categoryService.CategoryRepo.GetCategoryById(id)
	if categoryErr != nil {
		return nil, pkg_types.NewServerError(
			"error in get category from database",
			"CategoryService.GetCategoryById",
			categoryErr.Error(),
		)
	}
	if category == nil {
		return nil, pkg_types.NewClientError(http.StatusNotFound, "category not found by this id")
	}
	return category, nil
}

func (categoryService CategoryService) DeleteCategoryById(id uint) error {
	category, categoryErr := categoryService.GetCategoryById(id)
	if categoryErr != nil {
		return categoryErr
	}
	deleteCategoryErr := categoryService.CategoryRepo.DeleteCategoryById(category.ID)
	if deleteCategoryErr != nil {
		return pkg_types.NewServerError(
			"error in deleting category from database",
			"CategoryService.DeleteCategoryById",
			deleteCategoryErr.Error(),
		)
	}
	return nil
}
