package handlers

import (
	"backend/middleware"
	"backend/models"
	"backend/services"
	"encoding/json"
	"log"
	"net/http"
)

func GetPortfolioUser(w http.ResponseWriter, r *http.Request) {
	user, err := services.GetPortfolioUser()
	if err != nil {
		log.Printf("Error retrieving portfolio user: %v", err)
		middleware.JSONErrorResponse(w, http.StatusNotFound, "Portfolio user not found", err)
		return
	}
	middleware.JSONResponse(w, http.StatusOK, true, "Portfolio user retrieved successfully", user)
}

func UpdatePortfolioUser(w http.ResponseWriter, r *http.Request) {
	var updatedUser models.PortfolioUser
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		log.Printf("Error decoding request body: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	if err := services.UpdatePortfolioUser(&updatedUser); err != nil {
		log.Printf("Error updating portfolio user: %v", err)
		middleware.JSONErrorResponse(w, http.StatusInternalServerError, "Failed to update portfolio user", err)
		return
	}
	middleware.JSONResponse(w, http.StatusOK, true, "Portfolio user updated successfully", updatedUser)
}