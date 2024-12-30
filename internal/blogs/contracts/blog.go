package blogs_contracts

type BlogRepositoryContract interface {
	CreateBlog()
	DeleteBlogById()
	GetBlogById()
	GetBlogs()
}
