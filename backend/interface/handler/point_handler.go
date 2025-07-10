package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/usecase/point"
)

type PointHandler struct {
	pointUsecase point.PointUsecase
}

func NewPointHandler(pointUsecase point.PointUsecase) *PointHandler {
	return &PointHandler{
		pointUsecase: pointUsecase,
	}
}

type BalanceResponse struct {
	UserID  int `json:"user_id"`
	Balance int `json:"balance"`
}

type TransactionResponse struct {
	ID          int       `json:"id"`
	Amount      int       `json:"amount"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (h *PointHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
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

	balance, err := h.pointUsecase.GetBalance(r.Context(), userID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := BalanceResponse{
		UserID:  userID,
		Balance: balance,
	}

	respondSuccess(w, response)
}

func (h *PointHandler) GetTransactionHistory(w http.ResponseWriter, r *http.Request) {
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

	transactions, err := h.pointUsecase.GetTransactionHistory(r.Context(), userID, limit)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := make([]TransactionResponse, 0)
	for _, tx := range transactions {
		response = append(response, TransactionResponse{
			ID:          tx.ID,
			Amount:      tx.Amount,
			Type:        string(tx.Type),
			Description: tx.Description,
			CreatedAt:   tx.CreatedAt,
		})
	}

	respondSuccess(w, response)
}