package blogs_repositories

import (
	"strings"

	blogs_entities "github.com/ladmakhi81/realtime-blogs/internal/blogs/entities"
	pkg_storage "github.com/ladmakhi81/realtime-blogs/pkg/storage"
)

type BlogRepository struct {
	Storage pkg_storage.Storage
}

func NewBlogRepository(storage pkg_storage.Storage) BlogRepository {
	return BlogRepository{Storage: storage}
}

func (blogRepository BlogRepository) CreateBlog(blog *blogs_entities.Blog) error {
	command := `
		INSERT INTO _blogs 
		(title, content, created_by_id, category_id, tags)
		VALUES 
		($1, $2, $3, $4, $5)
		RETURNING 
		id, created_at, updated_at, title, content, tags;
	`
	row := blogRepository.Storage.DB.QueryRow(
		command,
		blog.Title,
		blog.Content,
		blog.CreatedBy.ID,
		blog.Category.ID,
		strings.Join(blog.Tags, ","),
	)
	var tags string
	scanErr := row.Scan(
		&blog.ID,
		&blog.CreatedAt,
		&blog.UpdatedAt,
		&blog.Title,
		&blog.Content,
		&tags,
	)
	blog.Tags = strings.Split(tags, ",")
	return scanErr
}

func (blogRepository BlogRepository) DeleteBlogById() {}

func (blogRepository BlogRepository) GetBlogById() {}

func (blogRepository BlogRepository) GetBlogs() {}
