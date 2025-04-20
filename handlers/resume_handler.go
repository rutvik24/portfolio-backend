package handlers

import (
	"backend/middleware"
	"backend/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetResume(w http.ResponseWriter, r *http.Request) {
	resume, err := services.GetResume()
	if err != nil {
		log.Printf("Error retrieving resume: %v", err)
		middleware.JSONErrorResponse(w, http.StatusNotFound, "Resume not found", err)
		return
	}
	middleware.JSONResponse(w, http.StatusOK, true, "Resume retrieved successfully", resume)
}

func CreateResume(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		log.Printf("Error parsing multipart form: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid form data", err)
		return
	}

	// Retrieve the file and its header
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error retrieving file from form: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "File is required", err)
		return
	}
	defer file.Close()

	// Call the service to create the resume
	resume, err := services.CreateResume(file, fileHeader)
	if err != nil {
		log.Printf("Error creating resume: %v", err)
		middleware.JSONErrorResponse(w, http.StatusInternalServerError, "Failed to create resume", err)
		return
	}

	middleware.JSONResponse(w, http.StatusCreated, true, "Resume created successfully", resume)
}

func UpdateResume(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		log.Printf("Error parsing multipart form: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid form data", err)
		return
	}

	// Retrieve the file and its header
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error retrieving file from form: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "File is required", err)
		return
	}
	defer file.Close()

	// Call the service to update the resume
	updatedResume, err := services.UpdateResume(file, fileHeader)
	if err != nil {
		log.Printf("Error updating resume: %v", err)
		middleware.JSONErrorResponse(w, http.StatusInternalServerError, "Failed to update resume", err)
		return
	}

	middleware.JSONResponse(w, http.StatusOK, true, "Resume updated successfully", updatedResume)
}

func DeleteResume(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Invalid resume ID: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid resume ID", err)
		return
	}
	if err := services.DeleteResume(uint(id)); err != nil {
		log.Printf("Error deleting resume: %v", err)
		middleware.JSONErrorResponse(w, http.StatusInternalServerError, "Failed to delete resume", err)
		return
	}
	middleware.JSONResponse(w, http.StatusNoContent, true, "Resume deleted successfully", nil)
}