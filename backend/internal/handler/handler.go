package handler

import (
	"encoding/json"
	"net/http"

	"github.com/nuchhbs/my-wealth/backend/internal/service"
)

type Handler struct {
	svc *service.FinanceService
}

func NewHandler(svc *service.FinanceService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	txns, err := h.svc.GetTransactions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, txns)
}

func (h *Handler) GetSummary(w http.ResponseWriter, r *http.Request) {
	summary, err := h.svc.GetSummary()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, summary)
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
