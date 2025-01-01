package blogs_types

import blogs_entities "github.com/ladmakhi81/realtime-blogs/internal/blogs/entities"

type GetBlogDetailResponse struct {
	Blog *blogs_entities.Blog `json:"blog"`
}

func NewGetBlogDetailResponse(blog *blogs_entities.Blog) GetBlogDetailResponse {
	return GetBlogDetailResponse{Blog: blog}
}
