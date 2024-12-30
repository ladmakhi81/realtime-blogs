package categories_handlers

import (
	"net/http"

	categories_services "github.com/ladmakhi81/realtime-blogs/internal/categories/services"
)

type CategoryHandler struct {
	CategoryService categories_services.CategoryService
}

func (categoryHandler CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	categoryHandler.CategoryService.CreateCategory()
}

func (categoryHandler CategoryHandler) DeleteCategoryById(w http.ResponseWriter, r *http.Request) {
	categoryHandler.CategoryService.DeleteCategoryById()
}

func (categoryHandler CategoryHandler) UpdateCategoryById(w http.ResponseWriter, r *http.Request) {
	categoryHandler.CategoryService.UpdateCategoryById()
}

func (categoryHandler CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categoryHandler.CategoryService.GetCategories()
}
