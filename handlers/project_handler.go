package handlers

import (
	"backend/middleware"
	"backend/models"
	"backend/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var validate = validator.New()

func GetProjects(w http.ResponseWriter, r *http.Request) {
	projects := services.GetAllProjects()
	middleware.JSONResponse(w, http.StatusOK, true, "Projects retrieved successfully", projects)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		log.Printf("Error decoding request body: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	// Validate the project struct
	if err := validate.Struct(project); err != nil {
		log.Printf("Validation error: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Validation failed", err)
		return
	}

	if err := services.CreateProject(&project); err != nil {
		log.Printf("Error creating project: %v", err)
		middleware.JSONErrorResponse(w, http.StatusInternalServerError, "Failed to create project", err)
		return
	}
	middleware.JSONResponse(w, http.StatusCreated, true, "Project created successfully", project)
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Invalid project ID: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid project ID", err)
		return
	}
	project, err := services.GetProjectByID(uint(id))
	if err != nil {
		log.Printf("Error retrieving project: %v", err)
		middleware.JSONErrorResponse(w, http.StatusNotFound, "Project not found", err)
		return
	}
	middleware.JSONResponse(w, http.StatusOK, true, "Project retrieved successfully", project)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Invalid project ID: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid project ID", err)
		return
	}
	var updatedProject models.Project
	if err := json.NewDecoder(r.Body).Decode(&updatedProject); err != nil {
		log.Printf("Error decoding request body: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	if err := services.UpdateProject(uint(id), &updatedProject); err != nil {
		log.Printf("Error updating project: %v", err)
		middleware.JSONErrorResponse(w, http.StatusInternalServerError, "Failed to update project", err)
		return
	}
	middleware.JSONResponse(w, http.StatusOK, true, "Project updated successfully", updatedProject)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Invalid project ID: %v", err)
		middleware.JSONErrorResponse(w, http.StatusBadRequest, "Invalid project ID", err)
		return
	}
	if err := services.DeleteProject(uint(id)); err != nil {
		log.Printf("Error deleting project: %v", err)
		middleware.JSONErrorResponse(w, http.StatusInternalServerError, "Failed to delete project", err)
		return
	}
	middleware.JSONResponse(w, http.StatusNoContent, true, "Project deleted successfully", nil)
}