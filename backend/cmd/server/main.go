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
	credFile := os.Getenv("GOOGLE_CREDENTIALS_FILE")

	sheetsClient, err := sheets.NewClient(credFile, sheetID)
	if err != nil {
		log.Fatalf("failed to init sheets client: %v", err)
	}

	svc := service.NewFinanceService(sheetsClient)
	h := handler.NewHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
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
