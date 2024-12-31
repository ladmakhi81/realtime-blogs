package categories_repositories

import (
	"database/sql"

	categories_entities "github.com/ladmakhi81/realtime-blogs/internal/categories/entities"
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

func (categoryRepository CategoryRepository) DeleteCategoryById() {}

func (categoryRepository CategoryRepository) GetCategories() {}

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
