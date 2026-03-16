package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"user-service/internal/handlers"
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
}
