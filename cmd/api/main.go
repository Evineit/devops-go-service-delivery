package main

import (
	"log"
	"net/http"
	"os"

	"user-service/internal/handlers"
	"user-service/internal/middleware"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/user/profile", handlers.HandleUser)
	http.HandleFunc("/users", handlers.HandleUsers)
	http.HandleFunc("/health", handlers.HandleHealth)
	http.HandleFunc("/metrics", handlers.HandleMetrics)

	// Read the port from environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server is running on http://localhost:" + port)
	// Wrap the default mux with the logging middleware
	log.Fatal(http.ListenAndServe(":"+port, middleware.LogRequest(http.DefaultServeMux)))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the User Profile API"))
}
