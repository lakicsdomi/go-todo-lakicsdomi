package routes_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-todo/controllers"
	"go-todo/routes"

	_ "github.com/go-sql-driver/mysql" // Required for the dummy database connection
)

// TestInitRoutes verifies that the application's routes are initialized correctly.
func TestInitRoutes(t *testing.T) {
	// Use a dummy database connection to avoid nil pointer issues in controllers.
	dummyDB, _ := sql.Open("mysql", "dummy:dummy@tcp(127.0.0.1:3306)/dummy")
	controllers.SetDatabase(dummyDB)

	router := routes.Init()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	// A non-404 response confirms that the route exists and reached the controller.
	if status := rr.Code; status == http.StatusNotFound {
		t.Errorf("Handler returned 404 Not Found, but the root route should exist")
	}

	req404, err := http.NewRequest("GET", "/invalid-route", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	rr404 := httptest.NewRecorder()

	router.ServeHTTP(rr404, req404)

	if status := rr404.Code; status != http.StatusNotFound {
		t.Errorf("Expected 404 for invalid route, got %v", status)
	}
}
