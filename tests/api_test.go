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

func TestHealthEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	handlers.HandleHealth(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	var body map[string]string
	if json.NewDecoder(resp.Body).Decode(&body) != nil {
		t.Fatal("Failed to decode JSON body for /health")
	}
	if body["status"] != "ok" {
		t.Errorf("Expected health.status = 'ok', got %v", body["status"])
	}
}

func TestListUsers(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()
	handlers.HandleUsers(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	var users []store.User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		t.Fatal("Failed to decode users list")
	}
	if len(users) == 0 {
		t.Error("Expected at least one user")
	}
}

func TestCreateUser(t *testing.T) {
	payload := `{"name": "John Doe", "email": "john.doe@example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.HandleUsers(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201 Created, got %d", resp.StatusCode)
	}

	var user store.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		t.Fatal("Failed to decode created user")
	}
	if user.Name != "John Doe" || user.Email != "john.doe@example.com" {
		t.Errorf("Created user data mismatch: %+v", user)
	}
	if user.Id == 0 {
		t.Error("Expected non-zero user ID")
	}
}

func TestCreateUserMalformedJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name": "John Doe", "email": "`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.HandleUsers(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request, got %d", resp.StatusCode)
	}
}

func TestCreateUserMissingFields(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name": "John Doe"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.HandleUsers(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request for missing fields, got %d", resp.StatusCode)
	}
}

func TestGetUserProfile(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/user/profile?clientId=1", nil)
	w := httptest.NewRecorder()
	handlers.HandleUser(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	var user store.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		t.Fatal("Failed to decode user profile")
	}
	if user.Id != 1 {
		t.Errorf("Expected user ID 1, got %d", user.Id)
	}
}

func TestGetUserProfileNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/user/profile?clientId=999", nil)
	w := httptest.NewRecorder()
	handlers.HandleUser(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status 404 Not Found, got %d", resp.StatusCode)
	}
}

func TestGetUserProfileInvalidID(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/user/profile?clientId=abc", nil)
	w := httptest.NewRecorder()
	handlers.HandleUser(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request, got %d", resp.StatusCode)
	}
}

func TestUpdateUserProfile(t *testing.T) {
	payload := `{"name": "Updated Name", "email": "updated@example.com"}`
	req := httptest.NewRequest(http.MethodPatch, "/user/profile?clientId=1", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.HandleUser(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	var user store.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		t.Fatal("Failed to decode updated user")
	}
	if user.Id != 1 {
		t.Errorf("Expected user ID 1, got %d", user.Id)
	}
	if user.Name != "Updated Name" {
		t.Errorf("Expected name 'Updated Name', got %s", user.Name)
	}
	if user.Email != "updated@example.com" {
		t.Errorf("Expected email 'updated@example.com', got %s", user.Email)
	}
}

func TestUpdateUserProfileNotFound(t *testing.T) {
	payload := `{"name": "Nope", "email": "nope@example.com"}`
	req := httptest.NewRequest(http.MethodPatch, "/user/profile?clientId=999", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.HandleUser(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status 404 Not Found, got %d", resp.StatusCode)
	}
}

func TestMetrics(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	handler := middleware.LogRequest(http.HandlerFunc(handlers.HandleHealth))
	handler.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/metrics", nil)
	handlers.HandleMetrics(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK for /metrics, got %d", resp.StatusCode)
	}

	var body map[string]int64
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatal("Failed to decode JSON body for /metrics")
	}
	if body["requests"] < 1 {
		t.Errorf("Expected requests >= 1, got %d", body["requests"])
	}
}
