package auth_routers

import (
	"github.com/gorilla/mux"
	auth_handlers "github.com/ladmakhi81/realtime-blogs/internal/auth/handlers"
	pkg_utils "github.com/ladmakhi81/realtime-blogs/pkg/utils"
)

type AuthRouter struct {
	ApiRouter   *mux.Router
	AuthHandler *auth_handlers.AuthHandler
}

func NewAuthRouter(apiRouter *mux.Router, authHandler *auth_handlers.AuthHandler) AuthRouter {
	return AuthRouter{ApiRouter: apiRouter, AuthHandler: authHandler}
}

func (authRouter AuthRouter) Setup() {
	authApi := authRouter.ApiRouter.PathPrefix("/auth").Subrouter()

	authApi.HandleFunc(
		"/login",
		pkg_utils.ErrorInterceptor(authRouter.AuthHandler.Login),
	).Methods("post")

	authApi.HandleFunc(
		"/signup",
		pkg_utils.ErrorInterceptor(authRouter.AuthHandler.Signup),
	).Methods("post")

	authApi.HandleFunc(
		"/refresh-token",
		pkg_utils.ErrorInterceptor(authRouter.AuthHandler.RefreshToken),
	).Methods("patch")

	authApi.HandleFunc(
		"/forget-password",
		pkg_utils.ErrorInterceptor(authRouter.AuthHandler.ForgetPassword),
	).Methods("patch")

	authApi.HandleFunc(
		"/profile",
		pkg_utils.ErrorInterceptor(authRouter.AuthHandler.Profile),
	).Methods("get")
}
