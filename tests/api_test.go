package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"user-service/internal/handlers"
	"user-service/internal/middleware"
	"user-service/internal/store"
)

func TestHandleUser(t *testing.T) {
	// Test GET /health endpoint
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	handlers.HandleHealth(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}
	// Validate JSON body
	var healthBody map[string]string
	if json.NewDecoder(resp.Body).Decode(&healthBody) != nil {
		t.Errorf("Failed to decode JSON body for /health")
	}
	if healthBody["status"] != "ok" {
		t.Errorf("Expected health.status = 'ok', got %v", healthBody["status"])
	}

	// Test GET /users endpoint
	req = httptest.NewRequest(http.MethodGet, "/users", nil)
	w = httptest.NewRecorder()
	handlers.HandleUsers(w, req)

	resp = w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	// Test POST /users endpoint with valid payload
	payload := `{"name": "John Doe", "email": "john.doe@example.com"}`
	req = httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	handlers.HandleUsers(w, req)
	resp = w.Result()
	// It should return 201 Created for valid payload
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201 Created for valid payload, got %d", resp.StatusCode)
	}
	// Check response body for new user
	var newUser store.User
	if json.NewDecoder(resp.Body).Decode(&newUser) != nil {
		t.Errorf("Failed to decode response body for new user")
	}

	// Test POST /users endpoint with malformed JSON
	req = httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name": "John Doe", "email": "`))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	handlers.HandleUsers(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request for malformed JSON, got %d", resp.StatusCode)
	}

	// Test POST /users with missing required fields
	req = httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name": "John Doe"}`))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	handlers.HandleUsers(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request for missing fields, got %d", resp.StatusCode)
	}

	// Test metrics: ensure middleware increments counter
	// Call a handler wrapped with logging middleware
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/health", nil)
	handler := middleware.LogRequest(http.HandlerFunc(handlers.HandleHealth))
	handler.ServeHTTP(w, req)

	// Now call /metrics to read the counter
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/metrics", nil)
	handlers.HandleMetrics(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK for /metrics, got %d", resp.StatusCode)
	}
	var metricsBody map[string]int64
	if json.NewDecoder(resp.Body).Decode(&metricsBody) != nil {
		t.Errorf("Failed to decode JSON body for /metrics")
	}
	if metricsBody["requests"] < 1 {
		t.Errorf("Expected requests >= 1, got %d", metricsBody["requests"])
	}
}
