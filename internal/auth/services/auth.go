package auth_services

import (
	"net/http"

	auth_contracts "github.com/ladmakhi81/realtime-blogs/internal/auth/contracts"
	auth_types "github.com/ladmakhi81/realtime-blogs/internal/auth/types"
	users_contracts "github.com/ladmakhi81/realtime-blogs/internal/users/contracts"
	pkg_types "github.com/ladmakhi81/realtime-blogs/pkg/types"
)

type AuthService struct {
	TokenService auth_contracts.TokenServiceContractor
	UserService  users_contracts.UserServiceContract
}

func NewAuthService(
	tokenService auth_contracts.TokenServiceContractor,
	userService users_contracts.UserServiceContract,
) AuthService {
	return AuthService{
		TokenService: tokenService,
		UserService:  userService,
	}
}

func (authService AuthService) Login(reqBody auth_types.LoginReqBody) (*auth_types.LoginResponse, error) {
	user, findByEmailErr := authService.UserService.FindByEmailAndPassword(reqBody.Email, reqBody.Password)
	if findByEmailErr != nil {
		return nil, findByEmailErr
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
	duplicatedUser, duplicatedUserErr := authService.UserService.FindByEmail(reqBody.Email)
	if duplicatedUserErr != nil {
		return nil, duplicatedUserErr
	}
	if duplicatedUser != nil {
		return nil, pkg_types.NewClientError(
			http.StatusConflict,
			"user with this email address already exist",
		)
	}

	user, createUserErr := authService.UserService.CreateUser(reqBody.Email, reqBody.Password)
	if createUserErr != nil {
		return nil, createUserErr
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
