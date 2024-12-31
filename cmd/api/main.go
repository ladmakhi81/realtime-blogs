package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	auth_handlers "github.com/ladmakhi81/realtime-blogs/internal/auth/handlers"
	auth_repositories "github.com/ladmakhi81/realtime-blogs/internal/auth/repositories"
	auth_routers "github.com/ladmakhi81/realtime-blogs/internal/auth/routers"
	auth_services "github.com/ladmakhi81/realtime-blogs/internal/auth/services"
	users_repositories "github.com/ladmakhi81/realtime-blogs/internal/users/repositories"
	users_services "github.com/ladmakhi81/realtime-blogs/internal/users/services"
	pkg_storage "github.com/ladmakhi81/realtime-blogs/pkg/storage"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("The env files not loaded")
	}

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	// database
	dbStorage := pkg_storage.Storage{}
	dbStorage.Init()

	// repos
	userRepo := users_repositories.NewUserRepository(dbStorage)
	tokenRepo := auth_repositories.NewTokenRepository(dbStorage)

	// services
	passwordHashService := users_services.NewPasswordHashService()
	tokenService := auth_services.NewTokenService(tokenRepo)
	userService := users_services.NewUserService(userRepo, passwordHashService)
	authService := auth_services.NewAuthService(tokenService, userService)

	// handlers
	authHandler := auth_handlers.NewAuthHandler(&authService)

	// routers
	authRouter := auth_routers.NewAuthRouter(apiRouter, &authHandler)
	authRouter.Setup()

	listenErr := http.ListenAndServe(":8080", apiRouter)

	if listenErr != nil {
		log.Fatalln(listenErr)
	}

	fmt.Println("main function invoked ...")
}
