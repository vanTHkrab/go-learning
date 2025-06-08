package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-api-basic/app/models"   // Assuming you have a models package
	"go-api-basic/app/services" // Assuming you have a services package
)

// PostController handles post-related HTTP requests.
type PostController struct {
	BaseController
	PostService *services.PostService // Inject PostService dependency
}

// NewPostController creates a new PostController instance.
func NewPostController(postService *services.PostService) *PostController {
	return &PostController{
		PostService: postService,
	}
}

// GetPostByID handles GET /posts/{id}
func (pc *PostController) GetPostByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id") // Example: getting from query param
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pc.SendErrorResponse(w, http.StatusBadRequest, "Invalid post ID")
		return
	}

	post, err := pc.PostService.GetPostByID(id)
	if err != nil {
		pc.SendErrorResponse(w, http.StatusNotFound, "Post not found")
		return
	}

	pc.SendJSONResponse(w, http.StatusOK, post)
}

// CreatePost handles POST /posts
func (pc *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		pc.SendErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	createdPost, err := pc.PostService.CreatePost(&post)
	if err != nil {
		pc.SendErrorResponse(w, http.StatusInternalServerError, "Failed to create post")
		return
	}

	pc.SendJSONResponse(w, http.StatusCreated, createdPost)
}

// ... other post-related methods
