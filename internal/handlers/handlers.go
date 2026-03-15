package handlers

import (
	"encoding/json"
	"net/http"

	"user-service/internal/store"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetUser(w, r)
	case http.MethodPatch:
		UpdateUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAllUserProfiles(w, r)
	case http.MethodPost:
		createUserProfile(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getAllUserProfiles(w http.ResponseWriter, r *http.Request) {
	var profiles = store.GetAllUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profiles)
}

func createUserProfile(w http.ResponseWriter, r *http.Request) {
	var newProfile store.User
	if err := json.NewDecoder(r.Body).Decode(&newProfile); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if newProfile.Email == "" || newProfile.Name == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}
	store.CreateUser(newProfile)
	w.WriteHeader(http.StatusCreated)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var clientId = r.URL.Query().Get("clientId")
	clientProfile, ok := store.GetUserById(clientId)
	if !ok {
		http.Error(w, "Client profile not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	response := store.User{
		Email: clientProfile.Email,
		Id:    clientProfile.Id,
		Name:  clientProfile.Name,
	}
	json.NewEncoder(w).Encode(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var clientId = r.URL.Query().Get("clientId")
	clientProfile, ok := store.GetUserById(clientId)
	if !ok {
		http.Error(w, "User profile not found", http.StatusNotFound)
		return
	}

	// Decode the JSON payload into a struct
	var payloadData store.User
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Update the user profile with the new data
	if payloadData.Email != "" {
		clientProfile.Email = payloadData.Email
	}
	if payloadData.Name != "" {
		clientProfile.Name = payloadData.Name
	}
	store.UpdateUserById(clientId, clientProfile)

	w.WriteHeader(http.StatusOK)
}
