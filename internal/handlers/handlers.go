package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

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

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
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
	newProfile = store.CreateUser(newProfile)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProfile)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var clientId = r.URL.Query().Get("clientId")
	clientIdInt, err := strconv.Atoi(clientId)
	if err != nil {
		http.Error(w, "Invalid client ID", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	user, found := store.GetUserById(clientIdInt)
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var clientId = r.URL.Query().Get("clientId")
	clientIdInt, err := strconv.Atoi(clientId)
	if err != nil {
		http.Error(w, "Invalid client ID", http.StatusBadRequest)
		return
	}

	// Decode the JSON payload into a struct
	var payloadData store.User
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Update the user in the database
	updated := store.UpdateUserById(clientIdInt, payloadData)
	if !updated {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
