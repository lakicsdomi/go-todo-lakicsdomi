package tests

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"go-todo/config"
	"go-todo/controllers"
	"go-todo/routes"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load("../.env") // Load environment variables for testing
}

// Runs a full CRUD lifecycle test.
func TestControllersIntegration(t *testing.T) {
	testDB := config.Database()
	controllers.SetDatabase(testDB)
	router := routes.Init()

	// Clean up table for a clean test state
	_, _ = testDB.Exec("TRUNCATE TABLE todos")

	// 1. Test Add (Create)
	formData := url.Values{}
	formData.Set("item", "Test Todo Item")
	reqAdd, _ := http.NewRequest("POST", "/add", strings.NewReader(formData.Encode()))
	reqAdd.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rrAdd := httptest.NewRecorder()
	router.ServeHTTP(rrAdd, reqAdd)

	if rrAdd.Code != http.StatusMovedPermanently && rrAdd.Code != http.StatusFound {
		t.Errorf("Expected redirect after add, got %v", rrAdd.Code)
	}

	// 2. Test Show (Read)
	reqShow, _ := http.NewRequest("GET", "/", nil)
	rrShow := httptest.NewRecorder()
	router.ServeHTTP(rrShow, reqShow)

	if rrShow.Code != http.StatusOK {
		t.Errorf("Expected 200 OK for Show, got %v", rrShow.Code)
	}
	if !strings.Contains(rrShow.Body.String(), "Test Todo Item") {
		t.Errorf("Expected to find 'Test Todo Item' in response body")
	}

	// Fetch the ID of the inserted item dynamically
	var id string
	err := testDB.QueryRow("SELECT id FROM todos WHERE item = ?", "Test Todo Item").Scan(&id)
	if err != nil {
		t.Fatalf("Failed to fetch inserted id: %v", err)
	}

	// 3. Test Complete (Update)
	reqComplete, _ := http.NewRequest("GET", "/complete/"+id, nil)
	rrComplete := httptest.NewRecorder()
	router.ServeHTTP(rrComplete, reqComplete)

	if rrComplete.Code != http.StatusMovedPermanently && rrComplete.Code != http.StatusFound {
		t.Errorf("Expected redirect after complete, got %v", rrComplete.Code)
	}

	var completed bool
	_ = testDB.QueryRow("SELECT completed FROM todos WHERE id = ?", id).Scan(&completed)
	if !completed {
		t.Errorf("Expected item to be completed")
	}

	// 4. Test Delete (Delete)
	reqDelete, _ := http.NewRequest("GET", "/delete/"+id, nil)
	rrDelete := httptest.NewRecorder()
	router.ServeHTTP(rrDelete, reqDelete)

	if rrDelete.Code != http.StatusMovedPermanently && rrDelete.Code != http.StatusFound {
		t.Errorf("Expected redirect after delete, got %v", rrDelete.Code)
	}

	var count int
	_ = testDB.QueryRow("SELECT COUNT(*) FROM todos WHERE id = ?", id).Scan(&count)
	if count != 0 {
		t.Errorf("Expected item to be deleted")
	}

	// 5. Test Error Branches (Force DB closure to trigger error handling in controllers)
	_ = testDB.Close()

	reqShowErr, _ := http.NewRequest("GET", "/", nil)
	rrShowErr := httptest.NewRecorder()
	router.ServeHTTP(rrShowErr, reqShowErr)

	if rrShowErr.Code != http.StatusInternalServerError {
		t.Errorf("Expected 500 Internal Server Error when DB is closed, got %v", rrShowErr.Code)
	}

	// Add Error
	reqAddErr, _ := http.NewRequest("POST", "/add", strings.NewReader(formData.Encode()))
	reqAddErr.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rrAddErr := httptest.NewRecorder()
	router.ServeHTTP(rrAddErr, reqAddErr)

	// Delete Error
	reqDelErr, _ := http.NewRequest("GET", "/delete/999", nil)
	rrDelErr := httptest.NewRecorder()
	router.ServeHTTP(rrDelErr, reqDelErr)

	// Complete Error
	reqCompErr, _ := http.NewRequest("GET", "/complete/999", nil)
	rrCompErr := httptest.NewRecorder()
	router.ServeHTTP(rrCompErr, reqCompErr)
}
