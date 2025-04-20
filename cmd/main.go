package main

import (
	"log"
	"net/http"
	"strings"

	"backend/config"
	"backend/db"
	"backend/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database and auto-migrate models
	db.InitDB()
	db.AutoMigrate()

	// Seed default admin
	db.SeedDefaultAdmin()

	// Initialize router
	r := mux.NewRouter()

	// Register all routes
	r = routes.RegisterRoutes(r)

	// Configure CORS middleware
	allowedOrigins := config.GetEnv("CORS_ORIGINS", "")
	log.Printf("Allowed Origins: %s", allowedOrigins)
	originsOk := handlers.AllowedOrigins(strings.Split(allowedOrigins, ", "))
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "XMLHttpRequest", "Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	// Start server
	port := config.GetEnv("PORT", "8080")
	log.Printf("Server is running on port %s", port)
	log.Fatalf("Could not start server: %s", http.ListenAndServe(":"+ port, handlers.CORS(headersOk, originsOk, methodsOk)(r)))
}