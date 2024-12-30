package auth_handlers

import (
	"net/http"

	auth_services "github.com/ladmakhi81/realtime-blogs/internal/auth/services"
)

type AuthHandler struct {
	AuthService auth_services.AuthService
}

func (authHandler AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	authHandler.AuthService.Login()
}

func (authHandler AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	authHandler.AuthService.Signup()
}

func (authHandler AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	authHandler.AuthService.RefreshToken()
}

func (authHandler AuthHandler) ForgetPassword(w http.ResponseWriter, r *http.Request) {
	authHandler.AuthService.ForgetPassword()
}

func (authHandler AuthHandler) Profile(w http.ResponseWriter, r *http.Request) {
	authHandler.AuthService.Profile()
}
