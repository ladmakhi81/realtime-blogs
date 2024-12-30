package blogs_services

import blogs_contracts "github.com/ladmakhi81/realtime-blogs/internal/blogs/contracts"

type BlogService struct {
	BlogRepository blogs_contracts.BlogRepositoryContract
}

func (blogService BlogService) CreateBlog() {}

func (blogService BlogService) DeleteBlogById() {}

func (blogService BlogService) GetBlogs() {}

func (blogService BlogService) GetBlogById() {}
