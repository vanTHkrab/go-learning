package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go-api-basic/app/services"
	"go-api-basic/router"
)

func main() {
	// Initialize dependencies (e.g., database connection, services, repositories)
	// For simplicity, let's mock a service for now.
	userService := services.NewUserService() // Replace with actual service initialization
	postService := services.NewPostService()

	// Initialize the router
	r := mux.NewRouter()

	// Setup routes
	routes.SetupRoutes(r, userService, postService)

	// Start the server
	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
