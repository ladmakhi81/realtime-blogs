package auth_services

import (
	"net/http"

	auth_contracts "github.com/ladmakhi81/realtime-blogs/internal/auth/contracts"
	auth_types "github.com/ladmakhi81/realtime-blogs/internal/auth/types"
	users_contracts "github.com/ladmakhi81/realtime-blogs/internal/users/contracts"
	users_entities "github.com/ladmakhi81/realtime-blogs/internal/users/entities"
	pkg_types "github.com/ladmakhi81/realtime-blogs/pkg/types"
)

type AuthService struct {
	TokenRepository  auth_contracts.TokenRepositoryContract
	UserRepository   users_contracts.UserRepositoryContract
	TokenService     auth_contracts.TokenServiceContractor
	PasswordHashUtil auth_contracts.PasswordHashContract
}

func NewAuthService(
	tokenRepository auth_contracts.TokenRepositoryContract,
	userRepository users_contracts.UserRepositoryContract,
	tokenService auth_contracts.TokenServiceContractor,
	passwordHashUtil auth_contracts.PasswordHashContract,
) AuthService {
	return AuthService{
		TokenRepository:  tokenRepository,
		UserRepository:   userRepository,
		TokenService:     tokenService,
		PasswordHashUtil: passwordHashUtil,
	}
}

func (authService AuthService) Login(reqBody auth_types.LoginReqBody) (*auth_types.LoginResponse, error) {
	user, emailErr := authService.UserRepository.FindByEmail(reqBody.Email)
	if emailErr != nil {
		return nil, pkg_types.NewServerError(
			"error in finding user by email",
			"AuthService.Login.FindByEmail",
			emailErr.Error(),
		)
	}
	if user == nil {
		return nil, pkg_types.NewClientError(
			http.StatusNotFound,
			"unable to find user by this email address and password",
		)
	}
	if isValid := authService.PasswordHashUtil.CompareHashedText(reqBody.Password, user.Password); !isValid {
		return nil, pkg_types.NewClientError(
			http.StatusNotFound,
			"unable to find user by this email address and password",
		)
	}
	token, tokenErr := authService.TokenService.CreateToken(user)
	if tokenErr != nil {
		return nil, tokenErr
	}
	return &auth_types.LoginResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (authService AuthService) Signup(reqBody auth_types.SignupReqBody) (*auth_types.SignupResponse, error) {
	duplicatedUser, duplicatedUserErr := authService.UserRepository.FindByEmail(reqBody.Email)
	if duplicatedUserErr != nil {
		return nil, pkg_types.NewServerError(
			"error in finding user by email address",
			"AuthService.Signup.FindByEmail",
			duplicatedUserErr.Error(),
		)
	}
	if duplicatedUser != nil {
		return nil, pkg_types.NewClientError(
			http.StatusConflict,
			"user with this email address already exist",
		)
	}
	hashedPasswored, hashedPasswordErr := authService.PasswordHashUtil.HashText(reqBody.Password)
	if hashedPasswordErr != nil {
		return nil, pkg_types.NewServerError(
			"error in hashing password",
			"AuthService.Signup.HashText",
			hashedPasswordErr.Error(),
		)
	}
	user := users_entities.NewUser(reqBody.Email, hashedPasswored)
	if createUserErr := authService.UserRepository.CreateUser(user); createUserErr != nil {
		return nil, pkg_types.NewServerError(
			"error in creating user",
			"AuthService.Signup.CreateUser",
			createUserErr.Error(),
		)
	}
	token, tokenErr := authService.TokenService.CreateToken(user)
	if tokenErr != nil {
		return nil, tokenErr
	}
	return &auth_types.SignupResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (authService AuthService) RefreshToken() {}

func (authService AuthService) ForgetPassword() {}

func (authService AuthService) Profile() {}
