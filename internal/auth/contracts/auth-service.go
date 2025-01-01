package auth_contracts

import auth_types "github.com/ladmakhi81/realtime-blogs/internal/auth/types"

type AuthServiceContract interface {
	Login(reqBody auth_types.LoginReqBody) (*auth_types.LoginResponse, error)
	Signup(reqBody auth_types.SignupReqBody) (*auth_types.SignupResponse, error)
}
