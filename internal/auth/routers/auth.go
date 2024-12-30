package routers_auth

import (
	"github.com/gorilla/mux"
	auth_handlers "github.com/ladmakhi81/realtime-blogs/internal/auth/handlers"
)

type AuthRouter struct {
	ApiRouter   *mux.Router
	AuthHandler auth_handlers.AuthHandler
}

func (authRouter *AuthRouter) Setup() {
	authApi := authRouter.ApiRouter.PathPrefix("/auth").Subrouter()

	authApi.HandleFunc(
		"/login",
		authRouter.AuthHandler.Login,
	).Methods("post")

	authApi.HandleFunc(
		"/signup",
		authRouter.AuthHandler.Signup,
	).Methods("post")

	authApi.HandleFunc(
		"/refresh-token",
		authRouter.AuthHandler.RefreshToken,
	).Methods("patch")

	authApi.HandleFunc(
		"/forget-password",
		authRouter.AuthHandler.ForgetPassword,
	).Methods("patch")

	authApi.HandleFunc(
		"/profile",
		authRouter.AuthHandler.Profile,
	).Methods("get")
}
