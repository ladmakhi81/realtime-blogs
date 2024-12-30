package blogs_handlers

import "net/http"

type BlogHandler struct{}

func (blogHandler BlogHandler) CreateBlog(w http.ResponseWriter, r *http.Request) {}

func (blogHandler BlogHandler) DeleteBlogById(w http.ResponseWriter, r *http.Request) {}

func (blogHandler BlogHandler) GetBlogs(w http.ResponseWriter, r *http.Request) {}

func (blogHandler BlogHandler) GetBlogById(w http.ResponseWriter, r *http.Request) {}
