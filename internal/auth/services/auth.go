package auth_services

import (
	"net/http"

	auth_contracts "github.com/ladmakhi81/realtime-blogs/internal/auth/contracts"
	auth_entities "github.com/ladmakhi81/realtime-blogs/internal/auth/entities"
	auth_types "github.com/ladmakhi81/realtime-blogs/internal/auth/types"
	users_contracts "github.com/ladmakhi81/realtime-blogs/internal/users/contracts"
	users_entities "github.com/ladmakhi81/realtime-blogs/internal/users/entities"
	pkg_types "github.com/ladmakhi81/realtime-blogs/pkg/types"
	pkg_utils "github.com/ladmakhi81/realtime-blogs/pkg/utils"
)

type AuthService struct {
	TokenRepository auth_contracts.TokenRepositoryContract
	UserRepository  users_contracts.UserRepositoryContract
}

func NewAuthService(tokenRepository auth_contracts.TokenRepositoryContract, userRepository users_contracts.UserRepositoryContract) AuthService {
	return AuthService{
		TokenRepository: tokenRepository,
		UserRepository:  userRepository,
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
	if isValid := pkg_utils.CompareHashedText(reqBody.Password, user.Password); !isValid {
		return nil, pkg_types.NewClientError(
			http.StatusNotFound,
			"unable to find user by this email address and password",
		)
	}
	accessToken, accessTokenErr := pkg_utils.GenerateToken(user)
	if accessTokenErr != nil {
		return nil, pkg_types.NewServerError(
			"error in generating access token",
			"AuthService.Login.GenerateToken",
			accessTokenErr.Error(),
		)
	}
	refreshToken, refreshTokenErr := pkg_utils.GenerateRefreshToken()
	if refreshTokenErr != nil {
		return nil, pkg_types.NewServerError(
			"error in generating refresh token",
			"AuthService.Login.GenerateRefreshToken",
			refreshTokenErr.Error(),
		)
	}
	token := auth_entities.NewToken(accessToken, refreshToken, user)
	if createTokenErr := authService.TokenRepository.CreateToken(token); createTokenErr != nil {
		return nil, pkg_types.NewServerError(
			"error in saving token on database",
			"AuthService.Login.CreateToken",
			createTokenErr.Error(),
		)
	}
	return &auth_types.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
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
	user := users_entities.NewUser(reqBody.Email, reqBody.Password)
	if createUserErr := authService.UserRepository.CreateUser(user); createUserErr != nil {
		return nil, pkg_types.NewServerError(
			"error in creating user",
			"AuthService.Signup.CreateUser",
			createUserErr.Error(),
		)
	}
	accessToken, accessTokenErr := pkg_utils.GenerateToken(user)
	if accessTokenErr != nil {
		return nil, pkg_types.NewServerError(
			"error in generating access token",
			"AuthService.Signup.GenerateToken",
			accessTokenErr.Error(),
		)
	}
	refreshToken, refreshTokenErr := pkg_utils.GenerateRefreshToken()
	if refreshTokenErr != nil {
		return nil, pkg_types.NewServerError(
			"error in generating refresh token",
			"AuthService.Signup.GenerateRefreshToken",
			refreshTokenErr.Error(),
		)
	}
	token := auth_entities.NewToken(accessToken, refreshToken, user)
	if createTokenErr := authService.TokenRepository.CreateToken(token); createTokenErr != nil {
		return nil, pkg_types.NewServerError(
			"error in saving token on database",
			"AuthService.Signup.CreateToken",
			createTokenErr.Error(),
		)
	}
	return &auth_types.SignupResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (authService AuthService) RefreshToken() {}

func (authService AuthService) ForgetPassword() {}

func (authService AuthService) Profile() {}
