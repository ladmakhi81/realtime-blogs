package blogs_repositories

import (
	"database/sql"
	"strings"

	blogs_entities "github.com/ladmakhi81/realtime-blogs/internal/blogs/entities"
	categories_entities "github.com/ladmakhi81/realtime-blogs/internal/categories/entities"
	users_entities "github.com/ladmakhi81/realtime-blogs/internal/users/entities"
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

func (blogRepository BlogRepository) DeleteBlogById(id uint) error {
	command := `
		DELETE FROM _blogs WHERE id = $1
	`
	statement, prepareErr := blogRepository.Storage.DB.Prepare(command)
	defer statement.Close()
	if prepareErr != nil {
		return prepareErr
	}
	_, executeErr := statement.Exec(id)
	if executeErr != nil {
		return executeErr
	}
	return nil
}

func (blogRepository BlogRepository) GetBlogById(id uint) (*blogs_entities.Blog, error) {
	command := `
		SELECT 
		b.id, b.title, b.content, b.tags, b.created_at, b.updated_at,
		u.id, u.email, u.created_at, u.updated_at,
		c.id, c.title, c.created_at, c.updated_at
		FROM _blogs b
		INNER JOIN _users u ON b.created_by_id = u.id
		INNER JOIN _categories c ON b.category_id = c.id
		WHERE b.id = $1
	`
	row := blogRepository.Storage.DB.QueryRow(command, id)
	blog := new(blogs_entities.Blog)
	category := categories_entities.Category{}
	user := users_entities.User{}
	var tags string
	scanErr := row.Scan(
		&blog.ID,
		&blog.Title,
		&blog.Content,
		&tags,
		&blog.CreatedAt,
		&blog.UpdatedAt,
		&user.ID,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
		&category.ID,
		&category.Title,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, nil
		}
		return nil, scanErr
	}
	blog.Tags = strings.Split(tags, ",")
	blog.Category = category
	blog.CreatedBy = user
	return blog, nil
}

func (blogRepository BlogRepository) GetBlogs(page, limit uint) (*[]blogs_entities.Blog, error) {
	command := `
		SELECT
		b.id, b.title, b.content, b.tags, b.created_at, b.updated_at,
		c.id, c.title,
		u.id, u.email, u.first_name, u.last_name
		FROM _blogs b
		INNER JOIN _categories c ON c.id = b.category_id
		INNER JOIN _users u ON u.id = b.created_by_id
		ORDER BY b.id DESC
		LIMIT $1 OFFSET $2;
	`
	rows, queryErr := blogRepository.Storage.DB.Query(command, limit, page*limit)
	if queryErr != nil {
		return nil, queryErr
	}
	defer rows.Close()
	blogs := []blogs_entities.Blog{}
	for rows.Next() {
		blog := blogs_entities.Blog{}
		category := categories_entities.Category{}
		user := users_entities.User{}
		var tags string
		scanErr := rows.Scan(
			&blog.ID, &blog.Title, &blog.Content, &tags, &blog.CreatedAt, &blog.UpdatedAt,
			&category.ID, &category.Title,
			&user.ID, &user.Email, &user.FirstName, &user.LastName,
		)
		if scanErr != nil {
			return nil, scanErr
		}
		blog.CreatedBy = user
		blog.Category = category
		blogs = append(blogs, blog)
	}
	return &blogs, nil
}

func (blogRepository BlogRepository) GetBlogsCount() (uint, error) {
	command := `
		SELECT COUNT(*) as total_blogs FROM _blogs;
	`
	var count uint
	row := blogRepository.Storage.DB.QueryRow(command)
	scanErr := row.Scan(
		&count,
	)
	if scanErr != nil {
		return 0, scanErr
	}
	return count, nil
}
