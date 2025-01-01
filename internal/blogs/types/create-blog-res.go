package blogs_types

import blogs_entities "github.com/ladmakhi81/realtime-blogs/internal/blogs/entities"

type CreateBlogResponse struct {
	Blog *blogs_entities.Blog `json:"blog"`
}

func NewCreateBlogResponse(blog *blogs_entities.Blog) CreateBlogResponse {
	return CreateBlogResponse{Blog: blog}
}
