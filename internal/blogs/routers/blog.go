package blogs_routers

import (
	"github.com/gorilla/mux"
	blogs_handlers "github.com/ladmakhi81/realtime-blogs/internal/blogs/handlers"
)

type BlogRouter struct {
	ApiRouter   *mux.Router
	BlogHandler blogs_handlers.BlogHandler
}

func (blogRouter *BlogRouter) Setup() {
	blogApi := blogRouter.ApiRouter.PathPrefix("/blogs").Subrouter()

	blogApi.HandleFunc(
		"/",
		blogRouter.BlogHandler.CreateBlog,
	).Methods("post")

	blogApi.HandleFunc(
		"/{id}",
		blogRouter.BlogHandler.DeleteBlogById,
	).Methods("delete")

	blogApi.HandleFunc(
		"/{id}",
		blogRouter.BlogHandler.GetBlogById,
	).Methods("get")

	blogApi.HandleFunc(
		"/",
		blogRouter.BlogHandler.GetBlogs,
	).Methods("get")
}
