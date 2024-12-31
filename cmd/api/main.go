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
	auth_utils "github.com/ladmakhi81/realtime-blogs/internal/auth/utils"
	users_repositories "github.com/ladmakhi81/realtime-blogs/internal/users/repositories"
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

	// utils
	tokenGeneratorUtil := auth_utils.NewJwtTokenGenerator()
	passwordHashUtil := auth_utils.NewPasswordHashUtil()

	// services
	tokenService := auth_services.NewTokenService(tokenRepo, tokenGeneratorUtil)
	authService := auth_services.NewAuthService(tokenRepo, userRepo, tokenService, passwordHashUtil)

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
