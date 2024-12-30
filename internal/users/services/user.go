package users_services

import users_contracts "github.com/ladmakhi81/realtime-blogs/internal/users/contracts"

type UserService struct {
	UserRepo users_contracts.UserRepositoryContract
}

func (userService UserService) ChangePassword() {}

func (userService UserService) EditUser() {}
