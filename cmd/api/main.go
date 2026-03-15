package main

import (
	"log"
	"net/http"

	"user-service/internal/handlers"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/user/profile", handlers.HandleUser)
	http.HandleFunc("/users", handlers.HandleUsers)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the User Profile API"))
}