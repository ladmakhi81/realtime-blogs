package auth_handlers

import "net/http"

type AuthHandler struct{}

func (authHandler AuthHandler) Login(w http.ResponseWriter, r *http.Request) {}

func (authHandler AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {}

func (authHandler AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {}

func (authHandler AuthHandler) ForgetPassword(w http.ResponseWriter, r *http.Request) {}

func (authHandler AuthHandler) Profile(w http.ResponseWriter, r *http.Request) {}
