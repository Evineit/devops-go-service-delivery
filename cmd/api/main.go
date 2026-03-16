package main

import (
	"log"
	"net/http"
	"os"

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

	// Read the port from environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server is running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the User Profile API"))
}