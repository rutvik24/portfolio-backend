package services

import (
	"backend/db"
	"backend/models"
	"log"
)

func GetAllProjects() []models.Project {
	var projects []models.Project
	if err := db.DB.Find(&projects).Error; err != nil {
		log.Printf("Error retrieving projects: %v", err)
	}
	return projects
}

func GetProjectByID(id uint) (*models.Project, error) {
	var project models.Project
	if err := db.DB.First(&project, id).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func CreateProject(project *models.Project) error {
	if err := db.DB.Create(project).Error; err != nil {
		log.Printf("Error creating project in database: %v", err)
		return err
	}
	return nil
}

func UpdateProject(id uint, updatedProject *models.Project) error {
	var project models.Project
	if err := db.DB.First(&project, id).Error; err != nil {
		return err
	}
	if err := db.DB.Model(&project).Updates(updatedProject).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProject(id uint) error {
	if err := db.DB.Delete(&models.Project{}, id).Error; err != nil {
		return err
	}
	return nil
}