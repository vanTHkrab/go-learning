package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-api-basic/app/models"
	"go-api-basic/app/services"
)

// UserController handles user-related HTTP requests.
type UserController struct {
	BaseController
	UserService *services.UserService // Inject UserService dependency
}

// NewUserController creates a new UserController instance.
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// GetUserByID handles GET /users/{id}
func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL path (example using mux or gorilla/mux)
	// For simplicity, let's assume we get it from query param for now.
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		uc.SendErrorResponse(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := uc.UserService.GetUserByID(id)
	if err != nil {
		uc.SendErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}

	uc.SendJSONResponse(w, http.StatusOK, user)
}

// CreateUser handles POST /users
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		uc.SendErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	createdUser, err := uc.UserService.CreateUser(&user)
	if err != nil {
		uc.SendErrorResponse(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	uc.SendJSONResponse(w, http.StatusCreated, createdUser)
}

// ... other user-related methods (UpdateUser, DeleteUser, etc.)
