package users_contracts

import users_entities "github.com/ladmakhi81/realtime-blogs/internal/users/entities"

type UserServiceContract interface {
	FindByEmail(email string) (*users_entities.User, error)
	FindByEmailAndPassword(email string, password string) (*users_entities.User, error)
	CreateUser(email string, password string) (*users_entities.User, error)
	FindUserById(id uint) (*users_entities.User, error)
}
