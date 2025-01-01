package users_routers

import (
	"net/http"

	"github.com/gorilla/mux"
	users_handlers "github.com/ladmakhi81/realtime-blogs/internal/users/handlers"
	pkg_decorators "github.com/ladmakhi81/realtime-blogs/pkg/decorators"
)

type UserRouter struct {
	ApiRouter   *mux.Router
	UserHandler users_handlers.UserHandler
}

func NewUserRouter(
	apiRouter *mux.Router,
	userHandler users_handlers.UserHandler,
) UserRouter {
	return UserRouter{
		ApiRouter:   apiRouter,
		UserHandler: userHandler,
	}
}

func (userRouter *UserRouter) Setup() {
	userApi := userRouter.ApiRouter.PathPrefix("/users").Subrouter()

	userApi.HandleFunc(
		"/edit-user",
		pkg_decorators.ApiErrorDecorator(
			pkg_decorators.ApiAuthDecorator(
				userRouter.UserHandler.EditUser,
			),
		),
	).Methods(http.MethodPatch)

	userApi.HandleFunc(
		"/upload-profile",
		pkg_decorators.ApiErrorDecorator(
			pkg_decorators.ApiAuthDecorator(
				userRouter.UserHandler.UploadProfile,
			),
		),
	).Methods(http.MethodPatch)
}
