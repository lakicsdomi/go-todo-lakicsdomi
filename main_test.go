package main

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var originalPort string

// Runs the main function but forces it to exit immediately
// by providing an invalid port, allowing us to cover the initialization logic.
func TestMainExecution(t *testing.T) {
	// Ensure environment variables are loaded
	err := godotenv.Load(".env")
	if err != nil {
		t.Logf("Error loading .env file: %v", err)
	}
	originalPort = os.Getenv("PORT")
	defer func() {
		if os.Setenv("PORT", originalPort) != nil { // Restore original PORT after the test
			t.Logf("Error restoring original PORT: %v", err)
		}
	}()
	// Set an invalid port to make http.ListenAndServe fail immediately.
	// This prevents main() from blocking the test runner forever.
	_ = os.Setenv("PORT", "-1")

	// Call the main function directly
	main()
}
