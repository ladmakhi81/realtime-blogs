package blogs_routers

import (
	"net/http"

	"github.com/gorilla/mux"
	blogs_handlers "github.com/ladmakhi81/realtime-blogs/internal/blogs/handlers"
	pkg_decorators "github.com/ladmakhi81/realtime-blogs/pkg/decorators"
)

type BlogRouter struct {
	ApiRouter   *mux.Router
	BlogHandler blogs_handlers.BlogHandler
}

func NewBlogRouter(
	apiRouter *mux.Router,
	blogHandler blogs_handlers.BlogHandler,
) BlogRouter {
	return BlogRouter{
		ApiRouter:   apiRouter,
		BlogHandler: blogHandler,
	}
}

func (blogRouter *BlogRouter) Setup() {
	blogApi := blogRouter.ApiRouter.PathPrefix("/blogs").Subrouter()

	blogApi.HandleFunc(
		"",
		pkg_decorators.ApiErrorDecorator(
			pkg_decorators.ApiAuthDecorator(
				blogRouter.BlogHandler.CreateBlog,
			),
		),
	).Methods(http.MethodPost)

	blogApi.HandleFunc(
		"/{id}",
		blogRouter.BlogHandler.DeleteBlogById,
	).Methods("delete")

	blogApi.HandleFunc(
		"/{id}",
		pkg_decorators.ApiErrorDecorator(
			pkg_decorators.ApiAuthDecorator(
				blogRouter.BlogHandler.GetBlogById,
			),
		),
	).Methods(http.MethodGet)

	blogApi.HandleFunc(
		"/",
		blogRouter.BlogHandler.GetBlogs,
	).Methods("get")
}
