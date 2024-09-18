package handlers

import (
	"encoding/json"
	"instagram-api/models"
	"instagram-api/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserRepo *repository.UserRepository
}

func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{UserRepo: userRepo}
}

// CreateUser handles the creation of a new user
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := user.HashPassword(); err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	if err := h.UserRepo.CreateUser(&user); err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GetUser handles retrieving a user by ID
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	user, err := h.UserRepo.GetUser(id)
	if err != nil {
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(user)
}
