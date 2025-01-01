package users_repositories

import (
	"database/sql"

	users_entities "github.com/ladmakhi81/realtime-blogs/internal/users/entities"
	pkg_storage "github.com/ladmakhi81/realtime-blogs/pkg/storage"
)

type UserRepository struct {
	Storage pkg_storage.Storage
}

func NewUserRepository(storage pkg_storage.Storage) UserRepository {
	return UserRepository{Storage: storage}
}

func (userRepo UserRepository) CreateUser(user *users_entities.User) error {
	command := `
		INSERT INTO 
		_users
		(email, password) VALUES ($1, $2)
		RETURNING id, created_at, updated_at, 
		email, password, profile_url, first_name, last_name;
	`
	row := userRepo.Storage.DB.QueryRow(command, user.Email, user.Password)
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
		SELECT * FROM _users WHERE email=$1 LIMIT 1;
	`
	user := new(users_entities.User)
	row := userRepo.Storage.DB.QueryRow(command, email)
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

func (userRepo UserRepository) FindUserById(id uint) (*users_entities.User, error) {
	command := `
		SELECT * FROM _users WHERE id = $1 LIMIT 1;
	`
	user := new(users_entities.User)
	row := userRepo.Storage.DB.QueryRow(command, id)
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
		return nil, scanErr
	}
	return user, nil
}

func (userRepo UserRepository) UpdateUserById(user *users_entities.User) error {
	command := `
		UPDATE _users SET
		first_name = $1,
		last_name = $2,
		profile_url = $3
		WHERE id = $4
	`
	statement, prepareErr := userRepo.Storage.DB.Prepare(command)
	if prepareErr != nil {
		return prepareErr
	}
	defer statement.Close()
	if _, executeErr := statement.Exec(
		user.FirstName,
		user.LastName,
		user.ProfileURL,
		user.ID,
	); executeErr != nil {
		return executeErr
	}
	return nil
}
