package categories_routers

import (
	"net/http"

	"github.com/gorilla/mux"
	categories_handlers "github.com/ladmakhi81/realtime-blogs/internal/categories/handlers"
	pkg_decorators "github.com/ladmakhi81/realtime-blogs/pkg/decorators"
)

type CategoryRouter struct {
	ApiRouter       *mux.Router
	CategoryHandler categories_handlers.CategoryHandler
}

func NewCategoryRouter(
	apiRouter *mux.Router,
	categoryHandler categories_handlers.CategoryHandler,
) CategoryRouter {
	return CategoryRouter{
		ApiRouter:       apiRouter,
		CategoryHandler: categoryHandler,
	}
}

func (categoryRouter CategoryRouter) Setup() {
	categoryApi := categoryRouter.ApiRouter.PathPrefix("/categories").Subrouter()

	categoryApi.HandleFunc(
		"",
		pkg_decorators.ApiErrorDecorator(
			pkg_decorators.ApiAuthDecorator(
				categoryRouter.CategoryHandler.CreateCategory,
			),
		),
	).Methods(http.MethodPost)

	categoryApi.HandleFunc(
		"/{id}",
		pkg_decorators.ApiErrorDecorator(
			pkg_decorators.ApiAuthDecorator(
				categoryRouter.CategoryHandler.DeleteCategoryById,
			),
		),
	).Methods(http.MethodDelete)

	categoryApi.HandleFunc(
		"/{id}",
		categoryRouter.CategoryHandler.UpdateCategoryById,
	).Methods(http.MethodPatch)

	categoryApi.HandleFunc(
		"",
		pkg_decorators.ApiErrorDecorator(
			pkg_decorators.ApiAuthDecorator(
				categoryRouter.CategoryHandler.GetCategories,
			),
		),
	).Methods(http.MethodGet)
}
