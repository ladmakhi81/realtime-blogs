package users_handlers

import "net/http"

type UserHandler struct{}

func (userHandler UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {}

func (userHandler UserHandler) EditUser(w http.ResponseWriter, r *http.Request) {}
