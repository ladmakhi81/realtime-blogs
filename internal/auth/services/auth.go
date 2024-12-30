package auth_services

import auth_contracts "github.com/ladmakhi81/realtime-blogs/internal/auth/contracts"

type AuthService struct {
	TokenRepository auth_contracts.TokenRepositoryContract
}

func (authService AuthService) Login() {}

func (authService AuthService) Signup() {}

func (authService AuthService) RefreshToken() {}

func (authService AuthService) ForgetPassword() {}

func (authService AuthService) Profile() {}
