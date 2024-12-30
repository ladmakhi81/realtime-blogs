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
	pkg_storage "github.com/ladmakhi81/realtime-blogs/pkg/storage"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("The env files not loaded")
	}

	router := mux.NewRouter()

	// database
	dbStorage := pkg_storage.Storage{}
	dbStorage.Init()

	// repos
	userRepo := users_repositories.NewUserRepository(dbStorage)
	tokenRepo := auth_repositories.NewTokenRepository(dbStorage)

	// services
	authService := auth_services.NewAuthService(tokenRepo, userRepo)

	// handlers
	authHandler := auth_handlers.NewAuthHandler(&authService)

	// routers
	authRouter := auth_routers.NewAuthRouter(router, &authHandler)
	authRouter.Setup()

	listenErr := http.ListenAndServe(":8080", router)

	if listenErr != nil {
		log.Fatalln(listenErr)
	}

	fmt.Println("main function invoked ...")
}
