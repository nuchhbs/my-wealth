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

// GET /api/statement — full financial statement (income, expenses, savings, summary)
func (h *Handler) GetStatement(w http.ResponseWriter, r *http.Request) {
	stmt, err := h.svc.GetStatement()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, stmt)
}

// GET /api/summary — summary only
func (h *Handler) GetSummary(w http.ResponseWriter, r *http.Request) {
	stmt, err := h.svc.GetStatement()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, stmt.Summary)
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(v)
}
