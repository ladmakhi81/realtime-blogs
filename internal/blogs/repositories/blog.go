package blogs_repositories

type BlogRepository struct{}

func (blogRepository BlogRepository) CreateBlog(
	title,
	content string,
	categoryId uint,
	tags []string,
	creatorId uint,
) {

}

func (blogRepository BlogRepository) DeleteBlogById() {}

func (blogRepository BlogRepository) GetBlogById() {}

func (blogRepository BlogRepository) GetBlogs() {}
