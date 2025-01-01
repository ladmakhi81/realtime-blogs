package blogs_contracts

import (
	blogs_entities "github.com/ladmakhi81/realtime-blogs/internal/blogs/entities"
	blogs_types "github.com/ladmakhi81/realtime-blogs/internal/blogs/types"
)

type BlogServiceContract interface {
	CreateBlog(
		reqBody blogs_types.CreateBlogReqBody,
		creatorId uint,
	) (*blogs_entities.Blog, error)
}
