package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nuchhbs/my-wealth/backend/internal/handler"
	"github.com/nuchhbs/my-wealth/backend/internal/service"
	"github.com/nuchhbs/my-wealth/backend/internal/sheets"
	"github.com/nuchhbs/my-wealth/backend/pkg/crypto"
)

func main() {
	sheetID := resolveSheetID()
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

// resolveSheetID prefers encrypted GOOGLE_SHEET_ID_ENC (needs SECRET_KEY),
// falls back to plaintext GOOGLE_SHEET_ID for local dev.
func resolveSheetID() string {
	enc := os.Getenv("GOOGLE_SHEET_ID_ENC")
	if enc != "" {
		secretKey := os.Getenv("SECRET_KEY")
		if secretKey == "" {
			log.Fatal("GOOGLE_SHEET_ID_ENC is set but SECRET_KEY is missing")
		}
		id, err := crypto.Decrypt(enc, secretKey)
		if err != nil {
			log.Fatalf("failed to decrypt GOOGLE_SHEET_ID_ENC: %v", err)
		}
		log.Println("using encrypted SHEET_ID ✓")
		return id
	}

	id := os.Getenv("GOOGLE_SHEET_ID")
	if id == "" {
		log.Fatal("either GOOGLE_SHEET_ID or GOOGLE_SHEET_ID_ENC must be set")
	}
	log.Println("warning: using plaintext GOOGLE_SHEET_ID (consider encrypting)")
	return id
}
