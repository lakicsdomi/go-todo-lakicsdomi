package tests

import (
	"testing"

	"go-todo/config"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load("../.env") // Load environment variables for testing
}

// Verifies that the database connects and initializes correctly.
func TestDatabaseSuccess(t *testing.T) {
	db := config.Database()
	if db == nil {
		t.Fatal("Expected database connection, got nil")
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing database connection: %v", err)
		}
	}()

	// Ping the database to ensure the connection is truly alive
	err := db.Ping()
	if err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}
}
