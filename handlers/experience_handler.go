package handlers

import (
	"backend/middleware"
	"backend/models"
	"backend/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetExperiences(w http.ResponseWriter, r *http.Request) {
	experiences := services.GetAllExperiences()
	middleware.JSONResponse(w, http.StatusOK, true, "Experiences retrieved successfully", experiences)
}

func CreateExperience(w http.ResponseWriter, r *http.Request) {
	var experience models.Experience
	if err := json.NewDecoder(r.Body).Decode(&experience); err != nil {
		log.Printf("Error decoding request body: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	// Set Present to true if EndDate is missing
	if experience.EndDate == "" {
		experience.Present = true
	}

	// Validate the experience struct
	if err := validate.Struct(experience); err != nil {
		log.Printf("Validation error: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Validation failed", err)
		return
	}

	if err := services.CreateExperience(&experience); err != nil {
		log.Printf("Error creating experience: %v", err)
		middleware.JSONErrorResponse(w, http.StatusInternalServerError, "Failed to create experience", err)
		return
	}
	middleware.JSONResponse(w, http.StatusCreated, true, "Experience created successfully", experience)
}

func GetExperience(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Invalid experience ID: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid experience ID", err)
		return
	}
	experience, err := services.GetExperienceByID(uint(id))
	if err != nil {
		log.Printf("Error retrieving experience: %v", err)
		middleware.JSONErrorResponse(w, http.StatusNotFound, "Experience not found", err)
		return
	}
	middleware.JSONResponse(w, http.StatusOK, true, "Experience retrieved successfully", experience)
}

func UpdateExperience(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Invalid experience ID: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid experience ID", err)
		return
	}
	var updatedExperience models.Experience
	if err := json.NewDecoder(r.Body).Decode(&updatedExperience); err != nil {
		log.Printf("Error decoding request body: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	if err := services.UpdateExperience(uint(id), &updatedExperience); err != nil {
		log.Printf("Error updating experience: %v", err)
		middleware.JSONErrorResponse(w, http.StatusInternalServerError, "Failed to update experience", err)
		return
	}
	middleware.JSONResponse(w, http.StatusOK, true, "Experience updated successfully", updatedExperience)
}

func DeleteExperience(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Invalid experience ID: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid experience ID", err)
		return
	}
	if err := services.DeleteExperience(uint(id)); err != nil {
		log.Printf("Error deleting experience: %v", err)
		middleware.JSONErrorResponse(w, http.StatusInternalServerError, "Failed to delete experience", err)
		return
	}
	middleware.JSONResponse(w, http.StatusNoContent, true, "Experience deleted successfully", nil)
}