package categories_handlers

import "net/http"

type CategoryHandler struct{}

func (categoryHandler CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {}

func (categoryHandler CategoryHandler) DeleteCategoryById(w http.ResponseWriter, r *http.Request) {}

func (categoryHandler CategoryHandler) UpdateCategoryById(w http.ResponseWriter, r *http.Request) {}

func (categoryHandler CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {}
