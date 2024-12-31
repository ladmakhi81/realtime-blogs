package users_repositories

import (
	"database/sql"

	users_entities "github.com/ladmakhi81/realtime-blogs/internal/users/entities"
	pkg_storage "github.com/ladmakhi81/realtime-blogs/pkg/storage"
)

type UserRepository struct {
	DBStorage pkg_storage.Storage
}

func NewUserRepository(storage pkg_storage.Storage) UserRepository {
	return UserRepository{DBStorage: storage}
}

func (userRepo UserRepository) CreateUser(user *users_entities.User) error {
	command := `
		INSERT INTO 
		"_users"
		("email", "password") VALUES ($1, $2)
		RETURNING "id", "created_at", "updated_at", 
		"email", "password", "profile_url", "first_name", "last_name";
	`
	row := userRepo.DBStorage.DB.QueryRow(command, user.Email, user.Password)
	scanErr := row.Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Email,
		&user.Password,
		&user.ProfileURL,
		&user.FirstName,
		&user.LastName,
	)
	if scanErr != nil {
		return scanErr
	}
	return nil
}

func (userRepo UserRepository) FindByEmail(email string) (*users_entities.User, error) {
	command := `
		SELECT * FROM "_users" WHERE "email"=$1 LIMIT 1;
	`
	user := new(users_entities.User)
	row := userRepo.DBStorage.DB.QueryRow(command, email)
	scanErr := row.Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Email,
		&user.Password,
		&user.ProfileURL,
		&user.FirstName,
		&user.LastName,
	)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, scanErr
		}
	}
	return user, nil
}
