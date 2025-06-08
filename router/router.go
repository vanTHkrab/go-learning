package routes

import (
	"github.com/gorilla/mux" // Example using gorilla/mux

	"go-api-basic/app/http/controllers" // Adjust the import path as per your project structure
	"go-api-basic/app/services"         // If using service layer
)

// SetupRoutes initializes all application routes.
func SetupRoutes(router *mux.Router, userService *services.UserService, postService *services.PostService) {
	// Initialize controllers
	userController := controllers.NewUserController(userService)
	postController := controllers.NewPostController(postService)

	// User Routes
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/", userController.CreateUser).Methods("POST")
	userRouter.HandleFunc("/{id}", userController.GetUserByID).Methods("GET")
	// ... more user routes

	// Post Routes
	postRouter := router.PathPrefix("/posts").Subrouter()
	postRouter.HandleFunc("/", postController.CreatePost).Methods("POST")
	postRouter.HandleFunc("/{id}", postController.GetPostByID).Methods("GET")
	// ... more post routes
}
