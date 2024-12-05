package testing

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/The-Pirateship/requests"
)

func TestPut(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("Expected PUT method, got %s", r.Method)
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Failed to read request body: %v", err)
		}

		expectedBody := `{"key":"value"}`
		if strings.TrimSpace(string(body)) != expectedBody {
			t.Errorf("Expected body: %s, got: %s", expectedBody, body)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"updated"}`))
	}))
	defer mockServer.Close()

	payload := map[string]string{"key": "value"}
	resp, err := requests.Put(mockServer.URL, payload)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedBody := `{"status":"updated"}`
	if strings.TrimSpace(string(resp.Body)) != expectedBody {
		t.Errorf("Expected body: %s, got: %s", expectedBody, resp.Body)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code: %d, got: %d", http.StatusOK, resp.StatusCode)
	}
}
