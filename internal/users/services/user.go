package users_services

import (
	"net/http"

	users_contracts "github.com/ladmakhi81/realtime-blogs/internal/users/contracts"
	users_entities "github.com/ladmakhi81/realtime-blogs/internal/users/entities"
	pkg_types "github.com/ladmakhi81/realtime-blogs/pkg/types"
)

type UserService struct {
	UserRepo            users_contracts.UserRepositoryContract
	PasswordHashService users_contracts.PasswordHashServiceContract
}

func NewUserService(
	userRepo users_contracts.UserRepositoryContract,
	passwordHashService users_contracts.PasswordHashServiceContract,
) UserService {
	return UserService{UserRepo: userRepo, PasswordHashService: passwordHashService}
}

func (userService UserService) FindByEmail(email string) (*users_entities.User, error) {
	user, emailErr := userService.UserRepo.FindByEmail(email)
	if emailErr != nil {
		return nil, pkg_types.NewServerError(
			"error in finding user by email",
			"UserService.FindByEmail",
			emailErr.Error(),
		)
	}
	return user, nil
}

func (userService UserService) CreateUser(email, password string) (*users_entities.User, error) {
	hashedPasswored, hashedPasswordErr := userService.PasswordHashService.HashText(password)
	if hashedPasswordErr != nil {
		return nil, pkg_types.NewServerError(
			"error in hashing password",
			"AuthService.Signup.HashText",
			hashedPasswordErr.Error(),
		)
	}
	user := users_entities.NewUser(email, hashedPasswored)
	if createUserErr := userService.UserRepo.CreateUser(user); createUserErr != nil {
		return nil, pkg_types.NewServerError(
			"error in creating user",
			"AuthService.Signup.CreateUser",
			createUserErr.Error(),
		)
	}
	return user, nil
}

func (userService UserService) FindByEmailAndPassword(email string, password string) (*users_entities.User, error) {
	user, findErr := userService.FindByEmail(email)
	if findErr != nil {
		return nil, findErr
	}
	if user == nil {
		return nil, pkg_types.NewClientError(
			http.StatusNotFound,
			"unable to find user by this email address and password",
		)
	}
	if isValid := userService.PasswordHashService.CompareHashedText(password, user.Password); !isValid {
		return nil, pkg_types.NewClientError(
			http.StatusNotFound,
			"unable to find user by this email address and password",
		)
	}
	return user, nil
}

func (userService UserService) FindUserById(id uint) (*users_entities.User, error) {
	user, findUserErr := userService.UserRepo.FindUserById(id)
	if findUserErr != nil {
		return nil, pkg_types.NewServerError(
			"error in find user by id",
			"UserService.FindUserById",
			findUserErr.Error(),
		)
	}
	return user, nil
}
