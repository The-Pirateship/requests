package testing

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/The-Pirateship/requests"
)

func TestGet(t *testing.T) {
	// Create a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/test" {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer mockServer.Close()

	// Call the Get function
	resp, err := requests.Get(mockServer.URL + "/test")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Validate response
	expectedBody := `{"message": "success"}`
	if strings.TrimSpace(string(resp.Body)) != expectedBody {
		t.Errorf("Expected body: %s, got: %s", expectedBody, resp.Body)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code: %d, got: %d", http.StatusOK, resp.StatusCode)
	}
}
