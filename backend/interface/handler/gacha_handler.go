package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/usecase/gacha"
)

type GachaHandler struct {
	gachaUsecase gacha.GachaUsecase
}

func NewGachaHandler(gachaUsecase gacha.GachaUsecase) *GachaHandler {
	return &GachaHandler{
		gachaUsecase: gachaUsecase,
	}
}

type ExecuteGachaRequest struct {
	UserID int `json:"user_id"`
}

type GachaResultResponse struct {
	ID           int       `json:"id"`
	ItemName     string    `json:"item_name"`
	Rarity       string    `json:"rarity"`
	PointsEarned int       `json:"points_earned"`
	CreatedAt    time.Time `json:"created_at"`
}

func (h *GachaHandler) ExecuteGacha(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req ExecuteGachaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.UserID <= 0 {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	result, err := h.gachaUsecase.ExecuteGacha(r.Context(), req.UserID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := GachaResultResponse{
		ID:           result.ID,
		ItemName:     result.ItemName,
		Rarity:       result.Rarity.String(),
		PointsEarned: result.PointsEarned,
		CreatedAt:    result.CreatedAt,
	}

	respondSuccess(w, response)
}

func (h *GachaHandler) GetGachaHistory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		respondError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID <= 0 {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	limitStr := r.URL.Query().Get("limit")
	limit := 20 // デフォルト値
	if limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	results, err := h.gachaUsecase.GetGachaHistory(r.Context(), userID, limit)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := make([]GachaResultResponse, 0)
	for _, result := range results {
		response = append(response, GachaResultResponse{
			ID:           result.ID,
			ItemName:     result.ItemName,
			Rarity:       result.Rarity.String(),
			PointsEarned: result.PointsEarned,
			CreatedAt:    result.CreatedAt,
		})
	}

	respondSuccess(w, response)
}