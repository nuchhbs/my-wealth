package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nuchhbs/my-wealth/backend/internal/handler"
	"github.com/nuchhbs/my-wealth/backend/internal/service"
	"github.com/nuchhbs/my-wealth/backend/internal/sheets"
	"github.com/nuchhbs/my-wealth/backend/pkg/crypto"
)

func main() {
	// Load .env from common locations
	for _, path := range []string{".env", "../.env", "../../.env"} {
		if err := godotenv.Load(path); err == nil {
			break
		}
	}

	sheetID := resolveSheetID()
	gid := os.Getenv("GOOGLE_SHEET_GID")

	sheetsClient := sheets.NewClient(sheetID, gid)
	svc := service.NewFinanceService(sheetsClient)
	h := handler.NewHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})
	mux.HandleFunc("/api/statement", h.GetStatement)
	mux.HandleFunc("/api/summary", h.GetSummary)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("server listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

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
	return id
}
