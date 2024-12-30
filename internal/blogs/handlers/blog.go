package blogs_handlers

import (
	"net/http"

	blogs_services "github.com/ladmakhi81/realtime-blogs/internal/blogs/services"
)

type BlogHandler struct {
	BlogService blogs_services.BlogService
}

func (blogHandler BlogHandler) CreateBlog(w http.ResponseWriter, r *http.Request) {
	blogHandler.BlogService.CreateBlog()
}

func (blogHandler BlogHandler) DeleteBlogById(w http.ResponseWriter, r *http.Request) {
	blogHandler.BlogService.DeleteBlogById()
}

func (blogHandler BlogHandler) GetBlogs(w http.ResponseWriter, r *http.Request) {
	blogHandler.BlogService.GetBlogs()
}

func (blogHandler BlogHandler) GetBlogById(w http.ResponseWriter, r *http.Request) {
	blogHandler.BlogService.GetBlogById()
}
