// CLI tool to encrypt a value with a secret key
// Usage: go run ./cmd/encrypt <value> <secret_key>
// Example: go run ./cmd/encrypt "my-sheet-id" "my-strong-secret"
package main

import (
	"fmt"
	"os"

	"github.com/nuchhbs/my-wealth/backend/pkg/crypto"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run ./cmd/encrypt <value> <secret_key>")
		os.Exit(1)
	}

	value := os.Args[1]
	secret := os.Args[2]

	encrypted, err := crypto.Encrypt(value, secret)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Encrypted value (put this in .env):")
	fmt.Printf("GOOGLE_SHEET_ID_ENC=%s\n", encrypted)
}
