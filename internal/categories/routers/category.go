package categories_routers

import (
	"github.com/gorilla/mux"
	categories_handlers "github.com/ladmakhi81/realtime-blogs/internal/categories/handlers"
)

type CategoryRouter struct {
	ApiRouter       *mux.Router
	CategoryHandler categories_handlers.CategoryHandler
}

func (categoryRouter CategoryRouter) Setup() {
	categoryApi := categoryRouter.ApiRouter.PathPrefix("/categories").Subrouter()

	categoryApi.HandleFunc(
		"/",
		categoryRouter.CategoryHandler.CreateCategory,
	).Methods("post")

	categoryApi.HandleFunc(
		"/{id}",
		categoryRouter.CategoryHandler.DeleteCategoryById,
	).Methods("delete")

	categoryApi.HandleFunc(
		"/{id}",
		categoryRouter.CategoryHandler.UpdateCategoryById,
	).Methods("patch")

	categoryApi.HandleFunc(
		"/",
		categoryRouter.CategoryHandler.GetCategories,
	).Methods("get")
}
