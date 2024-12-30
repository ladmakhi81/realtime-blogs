package users_handlers

import (
	"net/http"

	users_services "github.com/ladmakhi81/realtime-blogs/internal/users/services"
)

type UserHandler struct {
	UserService users_services.UserService
}

func (userHandler UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	userHandler.UserService.ChangePassword()
}

func (userHandler UserHandler) EditUser(w http.ResponseWriter, r *http.Request) {
	userHandler.UserService.EditUser()
}
