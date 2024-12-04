package testing

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/The-Pirateship/requests"
)

func TestDelete(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("Expected DELETE method, got %s", r.Method)
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer mockServer.Close()

	resp, err := requests.Delete(mockServer.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(resp.Body) != 0 {
		t.Errorf("Expected empty body, got: %s", resp.Body)
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected status code: %d, got: %d", http.StatusNoContent, resp.StatusCode)
	}
}
