package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"instagram-api/models"
	"instagram-api/repository"
	"log"
	"net/http"
	"strconv"
)

type PostHandler struct {
	PostRepo *repository.PostRepository
}

func NewPostHandler(postRepo *repository.PostRepository) *PostHandler {
	return &PostHandler{PostRepo: postRepo}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	if post.PostedTimestamp.IsZero() {
		http.Error(w, "Invalid timestamp format", http.StatusBadRequest)
		log.Printf("Invalid timestamp format: %v", post.PostedTimestamp)
		return
	}

	if err := h.PostRepo.CreatePost(&post); err != nil {
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		log.Printf("Error creating post: %v", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := h.PostRepo.GetPost(id)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) ListUserPosts(w http.ResponseWriter, r *http.Request) {
	userIDStr := mux.Vars(r)["userId"]
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	posts, err := h.PostRepo.ListUserPosts(userID)
	if err != nil {
		http.Error(w, "No posts found for user", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(posts)
}
