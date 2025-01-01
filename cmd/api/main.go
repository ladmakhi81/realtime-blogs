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
	blogs_handlers "github.com/ladmakhi81/realtime-blogs/internal/blogs/handlers"
	blogs_repositories "github.com/ladmakhi81/realtime-blogs/internal/blogs/repositories"
	blogs_routers "github.com/ladmakhi81/realtime-blogs/internal/blogs/routers"
	blogs_services "github.com/ladmakhi81/realtime-blogs/internal/blogs/services"
	categories_handlers "github.com/ladmakhi81/realtime-blogs/internal/categories/handlers"
	categories_repositories "github.com/ladmakhi81/realtime-blogs/internal/categories/repositories"
	categories_routers "github.com/ladmakhi81/realtime-blogs/internal/categories/routers"
	categories_services "github.com/ladmakhi81/realtime-blogs/internal/categories/services"
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
	categoryRepo := categories_repositories.NewCategoryRepository(dbStorage)
	blogRepo := blogs_repositories.NewBlogRepository(dbStorage)

	// services
	passwordHashService := users_services.NewPasswordHashService()
	tokenService := auth_services.NewTokenService(tokenRepo)
	userService := users_services.NewUserService(userRepo, passwordHashService)
	authService := auth_services.NewAuthService(tokenService, userService)
	categoryService := categories_services.NewCategoryService(categoryRepo, userService)
	blogService := blogs_services.NewBlogService(blogRepo, categoryService, userService)

	// handlers
	authHandler := auth_handlers.NewAuthHandler(authService)
	categoryHandler := categories_handlers.NewCategoryHandler(categoryService, userService)
	blogHandler := blogs_handlers.NewBlogHandler(blogService)

	// routers
	authRouter := auth_routers.NewAuthRouter(apiRouter, authHandler)
	categoryRouter := categories_routers.NewCategoryRouter(apiRouter, categoryHandler)
	blogRouter := blogs_routers.NewBlogRouter(apiRouter, blogHandler)

	authRouter.Setup()
	categoryRouter.Setup()
	blogRouter.Setup()

	listenErr := http.ListenAndServe(":8080", apiRouter)

	if listenErr != nil {
		log.Fatalln(listenErr)
	}

	fmt.Println("main function invoked ...")
}
