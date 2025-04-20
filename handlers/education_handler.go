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

func GetEducations(w http.ResponseWriter, r *http.Request) {
	educations := services.GetAllEducations()
	middleware.JSONResponse(w, http.StatusOK, true, "Educations retrieved successfully", educations)
}

func CreateEducation(w http.ResponseWriter, r *http.Request) {
	var education models.Education
	if err := json.NewDecoder(r.Body).Decode(&education); err != nil {

		log.Printf("Error decoding request body: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	// Validate the education struct
	if err := validate.Struct(education); err != nil {
		log.Printf("Validation error: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Validation failed", err)
		return
	}

	if err := services.CreateEducation(&education); err != nil {
		log.Printf("Error creating education: %v", err)
		middleware.JSONErrorResponse(w, http.StatusInternalServerError, "Failed to create education", err)
		return
	}
	log.Printf("Education created successfully: %v", education)
	middleware.JSONResponse(w, http.StatusCreated, true, "Education created successfully", education)
}

func GetEducation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Invalid education ID: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid education ID", err)
		return
	}
	education, err := services.GetEducationByID(uint(id))
	if err != nil {
		log.Printf("Error retrieving education: %v", err)
		middleware.JSONErrorResponse(w, http.StatusNotFound, "Education not found", err)
		return
	}
	middleware.JSONResponse(w, http.StatusOK, true, "Education retrieved successfully", education)
}

func UpdateEducation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Invalid education ID: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid education ID", err)
		return
	}
	var updatedEducation models.Education
	if err := json.NewDecoder(r.Body).Decode(&updatedEducation); err != nil {
		log.Printf("Error decoding request body: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	if err := services.UpdateEducation(uint(id), &updatedEducation); err != nil {
		log.Printf("Error updating education: %v", err)
		middleware.JSONErrorResponse(w, http.StatusInternalServerError, "Failed to update education", err)
		return
	}
	middleware.JSONResponse(w, http.StatusOK, true, "Education updated successfully", updatedEducation)
}

func DeleteEducation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Invalid education ID: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid education ID", err)
		return
	}
	if err := services.DeleteEducation(uint(id)); err != nil {
		log.Printf("Error deleting education: %v", err)
		middleware.JSONErrorResponse(w, http.StatusInternalServerError, "Failed to delete education", err)
		return
	}
	middleware.JSONResponse(w, http.StatusNoContent, true, "Education deleted successfully", nil)
}