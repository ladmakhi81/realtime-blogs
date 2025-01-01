package auth_routers

import (
	"net/http"

	"github.com/gorilla/mux"
	auth_handlers "github.com/ladmakhi81/realtime-blogs/internal/auth/handlers"
	pkg_decorators "github.com/ladmakhi81/realtime-blogs/pkg/decorators"
)

type AuthRouter struct {
	ApiRouter   *mux.Router
	AuthHandler auth_handlers.AuthHandler
}

func NewAuthRouter(
	apiRouter *mux.Router,
	authHandler auth_handlers.AuthHandler,
) AuthRouter {
	return AuthRouter{
		ApiRouter:   apiRouter,
		AuthHandler: authHandler,
	}
}

func (authRouter AuthRouter) Setup() {
	authApi := authRouter.ApiRouter.PathPrefix("/auth").Subrouter()

	authApi.HandleFunc(
		"/login",
		pkg_decorators.ApiErrorDecorator(authRouter.AuthHandler.Login),
	).Methods(http.MethodPost)

	authApi.HandleFunc(
		"/signup",
		pkg_decorators.ApiErrorDecorator(authRouter.AuthHandler.Signup),
	).Methods(http.MethodPost)
}
