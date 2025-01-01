package blogs_types

import blogs_entities "github.com/ladmakhi81/realtime-blogs/internal/blogs/entities"

type GetBlogsListResponse struct {
	Blogs       *[]blogs_entities.Blog `json:"blogs"`
	CurrentPage uint                   `json:"currentPage"`
	TotalPage   uint                   `json:"totalPage"`
	TotalBlogs  uint                   `json:"totalBlogs"`
}

func NewGetBlogsListResponse(
	blogs *[]blogs_entities.Blog,
	currentPage uint,
	totalPage uint,
	totalBlogs uint,
) GetBlogsListResponse {
	return GetBlogsListResponse{
		Blogs:       blogs,
		CurrentPage: currentPage,
		TotalPage:   totalPage,
		TotalBlogs:  totalBlogs,
	}
}
