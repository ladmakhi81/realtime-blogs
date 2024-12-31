package blogs_contracts

type BlogRepositoryContract interface {
	CreateBlog(
		title,
		content string,
		categoryId uint,
		tags []string,
		creatorId uint,
	)
	DeleteBlogById()
	GetBlogById()
	GetBlogs()
}
