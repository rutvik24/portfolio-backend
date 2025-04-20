package routes

import (
	"backend/handlers"
	"backend/middleware"

	"github.com/gorilla/mux"
)

func RegisterProjectRoutes(r *mux.Router) {
	r.HandleFunc("/projects", handlers.GetProjects).Methods("GET")
	r.HandleFunc("/projects", handlers.CreateProject).Methods("POST")
	r.HandleFunc("/projects/{id}", handlers.GetProject).Methods("GET")
	r.HandleFunc("/projects/{id}", handlers.UpdateProject).Methods("PUT")
	r.HandleFunc("/projects/{id}", handlers.DeleteProject).Methods("DELETE")
}

func RegisterAdminRoutes(r *mux.Router) {
	r.HandleFunc("/admins", handlers.CreateAdmin).Methods("POST")
	r.HandleFunc("/admins/authenticate", handlers.AuthenticateAdmin).Methods("POST")
	r.HandleFunc("/admins/{id}", handlers.DeleteAdmin).Methods("DELETE")
}

func RegisterExperienceRoutes(r *mux.Router) {
	r.HandleFunc("/experiences", handlers.GetExperiences).Methods("GET")
	r.HandleFunc("/experiences", handlers.CreateExperience).Methods("POST")
	r.HandleFunc("/experiences/{id}", handlers.GetExperience).Methods("GET")
	r.HandleFunc("/experiences/{id}", handlers.UpdateExperience).Methods("PUT")
	r.HandleFunc("/experiences/{id}", handlers.DeleteExperience).Methods("DELETE")
}

func RegisterEducationRoutes(r *mux.Router) {
	r.HandleFunc("/educations", handlers.GetEducations).Methods("GET")
	r.HandleFunc("/educations", handlers.CreateEducation).Methods("POST")
	r.HandleFunc("/educations/{id}", handlers.GetEducation).Methods("GET")
	r.HandleFunc("/educations/{id}", handlers.UpdateEducation).Methods("PUT")
	r.HandleFunc("/educations/{id}", handlers.DeleteEducation).Methods("DELETE")
}

func RegisterPortfolioUserRoutes(r *mux.Router) {
	r.HandleFunc("/portfolio-user", handlers.GetPortfolioUser).Methods("GET")
	r.HandleFunc("/portfolio-user", handlers.UpdatePortfolioUser).Methods("PUT")
}

func RegisterResumeRoutes(r *mux.Router) {
	r.HandleFunc("/resume", handlers.GetResume).Methods("GET")
	r.HandleFunc("/resume", handlers.CreateResume).Methods("POST")
	r.HandleFunc("/resume", handlers.UpdateResume).Methods("PUT")
	r.HandleFunc("/resume/{id}", handlers.DeleteResume).Methods("DELETE")
}

func RegisterRoutes(r *mux.Router) *mux.Router {
	// Apply JWT middleware to protected routes
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.JWTAuthMiddleware)
	
	// Register routes
	RegisterAdminRoutes(protected) // Admin routes are public
	RegisterProjectRoutes(protected)
	RegisterExperienceRoutes(protected)
	RegisterEducationRoutes(protected)
	RegisterPortfolioUserRoutes(protected)
	RegisterResumeRoutes(protected)

	return protected
}