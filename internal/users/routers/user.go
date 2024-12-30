package users_routers

import (
	"github.com/gorilla/mux"
	users_handlers "github.com/ladmakhi81/realtime-blogs/internal/users/handlers"
)

type UserRouter struct {
	ApiRouter   *mux.Router
	UserHandler users_handlers.UserHandler
}

func (userRouter *UserRouter) Setup() {
	userApi := userRouter.ApiRouter.PathPrefix("/users").Subrouter()

	userApi.HandleFunc(
		"/change-password",
		userRouter.UserHandler.ChangePassword,
	).Methods("patch")

	userApi.HandleFunc(
		"/edit-user",
		userRouter.UserHandler.EditUser,
	).Methods("patch")
}
