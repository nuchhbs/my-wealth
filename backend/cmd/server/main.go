package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nuchhbs/my-wealth/backend/internal/handler"
	"github.com/nuchhbs/my-wealth/backend/internal/service"
	"github.com/nuchhbs/my-wealth/backend/internal/sheets"
)

func main() {
	sheetID := os.Getenv("GOOGLE_SHEET_ID")
	if sheetID == "" {
		log.Fatal("GOOGLE_SHEET_ID is required")
	}
	gid := os.Getenv("GOOGLE_SHEET_GID") // optional, default = "0"

	sheetsClient := sheets.NewClient(sheetID, gid)
	svc := service.NewFinanceService(sheetsClient)
	h := handler.NewHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})
	mux.HandleFunc("/api/summary", h.GetSummary)
	mux.HandleFunc("/api/transactions", h.GetTransactions)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("server listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
