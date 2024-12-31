package categories_repositories

import (
	"database/sql"

	categories_entities "github.com/ladmakhi81/realtime-blogs/internal/categories/entities"
	users_entities "github.com/ladmakhi81/realtime-blogs/internal/users/entities"
	pkg_storage "github.com/ladmakhi81/realtime-blogs/pkg/storage"
)

type CategoryRepository struct {
	Storage pkg_storage.Storage
}

func NewCategoryRepository(
	storage pkg_storage.Storage,
) CategoryRepository {
	return CategoryRepository{Storage: storage}
}

func (categoryRepository CategoryRepository) CreateCategory(category *categories_entities.Category) error {
	command := `
		INSERT INTO "_categories"
		("title", "created_by_id") VALUES ($1, $2)
		RETURNING "id", "title", "created_at", "updated_at";
	`
	row := categoryRepository.Storage.DB.QueryRow(command, category.Title, category.CreatedBy.ID)
	scanErr := row.Scan(
		&category.ID,
		&category.Title,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	return scanErr
}

func (categoryRepository CategoryRepository) UpdateCategoryId() {}

func (categoryRepository CategoryRepository) DeleteCategoryById(id uint) error {
	command := `
		DELETE FROM _categories WHERE id = $1
	`
	statement, pErr := categoryRepository.Storage.DB.Prepare(command)
	if pErr != nil {
		return pErr
	}
	defer statement.Close()
	_, eErr := statement.Exec(id)
	if eErr != nil {
		return eErr
	}
	return nil
}

func (categoryRepository CategoryRepository) GetCategories(page, limit uint) (*[]categories_entities.Category, error) {
	command := `
		SELECT 
			c.id, c.title, c.created_at, c.updated_at,
			u.id, u.email, u.created_at, u.updated_at
		FROM _categories c
		INNER JOIN _users u ON u.id = c.created_by_id
		ORDER BY c.id DESC
		LIMIT $1 OFFSET $2
	`
	rows, queryErr := categoryRepository.Storage.DB.Query(command, limit, page*limit)
	defer rows.Close()
	if queryErr != nil {
		return nil, queryErr
	}
	categories := []categories_entities.Category{}
	for rows.Next() {
		category := categories_entities.Category{}
		createdBy := new(users_entities.User)
		scanErr := rows.Scan(
			&category.ID,
			&category.Title,
			&category.CreatedAt,
			&category.UpdatedAt,
			&createdBy.ID,
			&createdBy.Email,
			&createdBy.CreatedAt,
			&createdBy.UpdatedAt,
		)
		if scanErr != nil {
			return nil, scanErr
		}
		category.CreatedBy = createdBy
		categories = append(categories, category)
	}
	return &categories, nil
}

func (categoryRepository CategoryRepository) GetCategoryByTitle(title string) (*categories_entities.Category, error) {
	command := `
		SELECT "id", "title", "created_at", "updated_at" FROM "_categories" WHERE "title" = $1 LIMIT 1;
	`
	row := categoryRepository.Storage.DB.QueryRow(command, title)
	category := new(categories_entities.Category)
	scanErr := row.Scan(
		&category.ID,
		&category.Title,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, scanErr
		}
	}
	return category, nil
}

func (categoryRepository CategoryRepository) GetCategoryById(id uint) (*categories_entities.Category, error) {
	command := `
		SELECT 
			c.id, c.title, c.created_at, c.updated_at,
			u.id, u.email, u.created_at, u.updated_at
		FROM _categories c 
		INNER JOIN _users u ON u.id = c.created_by_id
		WHERE c.id = $1 
		LIMIT 1
	`
	row := categoryRepository.Storage.DB.QueryRow(command, id)
	category := new(categories_entities.Category)
	user := new(users_entities.User)

	scanErr := row.Scan(
		&category.ID,
		&category.Title,
		&category.CreatedAt,
		&category.UpdatedAt,
		&user.ID,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, scanErr
		}
	}
	category.CreatedBy = user
	return category, nil
}
