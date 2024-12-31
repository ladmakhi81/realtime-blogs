package auth_routers

import (
	"github.com/gorilla/mux"
	auth_handlers "github.com/ladmakhi81/realtime-blogs/internal/auth/handlers"
	pkg_decorators "github.com/ladmakhi81/realtime-blogs/pkg/decorators"
)

type AuthRouter struct {
	ApiRouter   *mux.Router
	AuthHandler auth_handlers.AuthHandler
}

func NewAuthRouter(apiRouter *mux.Router, authHandler auth_handlers.AuthHandler) AuthRouter {
	return AuthRouter{ApiRouter: apiRouter, AuthHandler: authHandler}
}

func (authRouter AuthRouter) Setup() {
	authApi := authRouter.ApiRouter.PathPrefix("/auth").Subrouter()

	authApi.HandleFunc(
		"/login",
		pkg_decorators.ApiErrorDecorator(authRouter.AuthHandler.Login),
	).Methods("post")

	authApi.HandleFunc(
		"/signup",
		pkg_decorators.ApiErrorDecorator(authRouter.AuthHandler.Signup),
	).Methods("post")

	authApi.HandleFunc(
		"/refresh-token",
		pkg_decorators.ApiErrorDecorator(authRouter.AuthHandler.RefreshToken),
	).Methods("patch")

	authApi.HandleFunc(
		"/forget-password",
		pkg_decorators.ApiErrorDecorator(authRouter.AuthHandler.ForgetPassword),
	).Methods("patch")

	authApi.HandleFunc(
		"/profile",
		pkg_decorators.ApiErrorDecorator(authRouter.AuthHandler.Profile),
	).Methods("get")
}
