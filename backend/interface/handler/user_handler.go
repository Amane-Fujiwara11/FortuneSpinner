package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/domain/model"
	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/domain/repository"
)

type UserHandler struct {
	userRepo repository.UserRepository
}

func NewUserHandler(userRepo repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

type UserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	user, err := model.NewUser(req.Name)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userRepo.Create(r.Context(), user); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := UserResponse{
		ID:   user.ID,
		Name: user.Name,
	}

	respondSuccess(w, response)
}

func (h *UserHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateUser(w, r)
	case http.MethodGet:
		h.GetUser(w, r)
	default:
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL path
	path := r.URL.Path
	
	// Handle both /api/users/123 and /api/users?id=123
	var userIDStr string
	if strings.HasPrefix(path, "/api/users/") {
		userIDStr = strings.TrimPrefix(path, "/api/users/")
	} else {
		userIDStr = r.URL.Query().Get("id")
	}
	
	if userIDStr == "" {
		respondError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.userRepo.FindByID(r.Context(), userID)
	if err != nil {
		respondError(w, http.StatusNotFound, "User not found")
		return
	}

	response := UserResponse{
		ID:   user.ID,
		Name: user.Name,
	}

	respondSuccess(w, response)
}