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
	users_handlers "github.com/ladmakhi81/realtime-blogs/internal/users/handlers"
	users_repositories "github.com/ladmakhi81/realtime-blogs/internal/users/repositories"
	users_routers "github.com/ladmakhi81/realtime-blogs/internal/users/routers"
	users_services "github.com/ladmakhi81/realtime-blogs/internal/users/services"
	pkg_storage "github.com/ladmakhi81/realtime-blogs/pkg/storage"
	pkg_utils "github.com/ladmakhi81/realtime-blogs/pkg/utils"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("The env files not loaded")
	}

	uploadedDirectory, uploadedDirectoryErr := pkg_utils.GetUploadedFileDirectory()

	if uploadedDirectoryErr != nil {
		log.Fatalln("unable to find root directory")
	}

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	router.PathPrefix("/uploads/").Handler(
		http.StripPrefix("/uploads/", http.FileServer(http.Dir(uploadedDirectory))),
	)

	// database
	dbStorage := pkg_storage.Storage{}
	dbStorage.Init()

	// repos
	userRepo := users_repositories.NewUserRepository(dbStorage)
	tokenRepo := auth_repositories.NewTokenRepository(dbStorage)
	categoryRepo := categories_repositories.NewCategoryRepository(dbStorage)
	blogRepo := blogs_repositories.NewBlogRepository(dbStorage)

	// services
	tokenService := auth_services.NewTokenService(tokenRepo)
	userService := users_services.NewUserService(userRepo)
	authService := auth_services.NewAuthService(tokenService, userService)
	categoryService := categories_services.NewCategoryService(categoryRepo, userService)
	blogService := blogs_services.NewBlogService(blogRepo, categoryService, userService)

	// handlers
	authHandler := auth_handlers.NewAuthHandler(authService)
	categoryHandler := categories_handlers.NewCategoryHandler(categoryService, userService)
	blogHandler := blogs_handlers.NewBlogHandler(blogService)
	userHandler := users_handlers.NewUserHandler(userService)

	// routers
	authRouter := auth_routers.NewAuthRouter(apiRouter, authHandler)
	categoryRouter := categories_routers.NewCategoryRouter(apiRouter, categoryHandler)
	blogRouter := blogs_routers.NewBlogRouter(apiRouter, blogHandler)
	userRouter := users_routers.NewUserRouter(apiRouter, userHandler)

	authRouter.Setup()
	categoryRouter.Setup()
	blogRouter.Setup()
	userRouter.Setup()

	listenErr := http.ListenAndServe(":8080", router)

	if listenErr != nil {
		log.Fatalln(listenErr)
	}

	fmt.Println("main function invoked ...")
}
