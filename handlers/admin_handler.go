package handlers

import (
	"backend/config"
	"backend/middleware"
	"backend/models"
	"backend/services"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

func GenerateJWT(adminID uint, role models.Role) (string, error) {
	secret := config.GetEnv("JWT_SECRET", "")
	if secret == "" {
		return "", errors.New("JWT secret not configured")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin_id": adminID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"role":     role,
	})

	return token.SignedString([]byte(secret))
}

func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	// Check if the requesting admin has the required role
	adminID := r.Context().Value("admin_id").(uint)
	requestingAdmin, err := services.GetAdminByID(adminID)
	if err != nil {
		log.Printf("Error fetching requesting admin: %v", err)
		middleware.JSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch admin", err)
		return
	}

	var admin models.Admin
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		log.Printf("Error decoding request body: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	// Role-based creation logic
	if requestingAdmin.Role == "admin" {
		middleware.JSONErrorResponse(w, http.StatusForbidden, "You do not have permission to create an admin", nil)
		return
	}

	if requestingAdmin.Role == "super_admin" && admin.Role == "super_super_admin" {
		middleware.JSONErrorResponse(w, http.StatusForbidden, "You do not have permission to create a super super admin", nil)
		return
	}

	if err := services.CreateAdmin(&admin); err != nil {
		log.Printf("Error creating admin: %v", err)
		middleware.JSONErrorResponse(w, http.StatusInternalServerError, "Failed to create admin", err)
		return
	}
	middleware.JSONResponse(w, http.StatusCreated, true, "Admin created successfully", nil)
}

func AuthenticateAdmin(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		log.Printf("Error decoding request body: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	admin, err := services.AuthenticateAdmin(credentials.Username, credentials.Password)
	if err != nil {
		log.Printf("Authentication failed: %v", err)
		middleware.JSONErrorResponse(w, http.StatusUnauthorized, "Invalid username or password", err)
		return
	}

	token, err := GenerateJWT(admin.ID, admin.Role)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		middleware.JSONErrorResponse(w, http.StatusInternalServerError, "Failed to generate token", err)
		return
	}

	
	// Update session token in the database
	if err := services.UpdateAdminSession(admin.ID, token); err != nil {
		log.Printf("Error updating session token: %v", err)
		middleware.JSONErrorResponse(w, http.StatusInternalServerError, "Failed to update session token", err)
		return
	}
	
	admin.Password = "" // Clear password before sending response
	admin.SessionToken = token

	middleware.JSONResponse(w, http.StatusOK, true, "Authentication successful", admin)
}

func DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	adminID := r.Context().Value("admin_id").(uint)
	requestingAdmin, err := services.GetAdminByID(adminID)
	if err != nil {
		log.Printf("Error fetching requesting admin: %v", err)
		middleware.JSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch admin", err)
		return
	}

	vars := mux.Vars(r)
	targetID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Invalid target admin ID: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid admin ID", err)
		return
	}

	if uint(targetID) == adminID {
		middleware.JSONErrorResponse(w, http.StatusForbidden, "You cannot delete yourself", nil)
		return
	}

	targetAdmin, err := services.GetAdminByID(uint(targetID))
	if err != nil {
		log.Printf("Error fetching target admin: %v", err)
		middleware.JSONErrorResponse(w, http.StatusNotFound, "Admin not found", err)
		return
	}

	if requestingAdmin.Role == "super_super_admin" || (requestingAdmin.Role == "super_admin" && targetAdmin.Role == "admin") {
		if err := services.DeleteAdmin(uint(targetID)); err != nil {
			log.Printf("Error deleting admin: %v", err)
			middleware.JSONErrorResponse(w, http.StatusInternalServerError, "Failed to delete admin", err)
			return
		}
		middleware.JSONResponse(w, http.StatusOK, true, "Admin deleted successfully", nil)
	} else {
		middleware.JSONErrorResponse(w, http.StatusForbidden, "You do not have permission to delete this admin", nil)
	}
}