package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"search-app/internal/auth"
	"search-app/internal/file"
	"search-app/internal/middleware"
	"search-app/internal/search"
	"search-app/internal/docs"
	"github.com/rs/cors"
	"search-app/database"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// CORS configuration
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow your frontend origin
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}, // Allowed HTTP methods
		AllowedHeaders:   []string{"Content-Type", "X-Clerk-User-Id", "Authorization"}, // Allowed headers
		AllowCredentials: true, // Allow cookies and credentials if needed
	})

	// Wrap the router with the CORS handler
	r.Use(corsHandler.Handler)

	// Public routes
	r.HandleFunc("/upload", file.UploadFileHandler).Methods("POST")
	r.HandleFunc("/search", search.SearchHandler).Methods("POST")
	r.HandleFunc("/search", search.SearchHandler).Methods("OPTIONS")
	r.HandleFunc("/docs/{id}", docs.DocHandler).Methods("GET")
	r.HandleFunc("/docsById/{id}", docs.DocByIdHandler).Methods("GET")

	// Protected routes (JWT authentication)
	protected := r.PathPrefix("/protected").Subrouter()
	protected.Use(middleware.AuthMiddleware) // Apply JWT authentication middleware to protected routes
	protected.HandleFunc("/profile", auth.ProfileHandler).Methods("GET")

	database.Init()
	// Start the server
	log.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
