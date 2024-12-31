package users_contracts

import users_entities "github.com/ladmakhi81/realtime-blogs/internal/users/entities"

type UserRepositoryContract interface {
	FindByEmail(email string) (*users_entities.User, error)
	CreateUser(user *users_entities.User) error
	FindUserById(id uint) (*users_entities.User, error)
}
