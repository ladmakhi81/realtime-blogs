package blogs_contracts

import blogs_entities "github.com/ladmakhi81/realtime-blogs/internal/blogs/entities"

type BlogRepositoryContract interface {
	CreateBlog(blog *blogs_entities.Blog) error
	DeleteBlogById()
	GetBlogById()
	GetBlogs()
}
