package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"plus-number/internal/db"
	"plus-number/internal/model"
)

type AddHandler struct {
	DB db.AddNumbersStore
}

func (h *AddHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req model.AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	sum, err := h.DB.AddNumbers(ctx, req.A, req.B)
	if err != nil {
		http.Error(w, "db error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := model.AddResponse{Result: sum}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
