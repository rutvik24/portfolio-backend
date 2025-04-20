package services

import (
	"backend/db"
	"backend/models"
	"log"
)

func GetAllExperiences() []models.Experience {
	var experiences []models.Experience
	if err := db.DB.Find(&experiences).Error; err != nil {
		log.Printf("Error retrieving experiences: %v", err)
	}
	return experiences
}

func GetExperienceByID(id uint) (*models.Experience, error) {
	var experience models.Experience
	if err := db.DB.First(&experience, id).Error; err != nil {
		return nil, err
	}
	return &experience, nil
}

func CreateExperience(experience *models.Experience) error {
	if err := db.DB.Create(experience).Error; err != nil {
		log.Printf("Error creating experience in database: %v", err)
		return err
	}
	return nil
}

func UpdateExperience(id uint, updatedExperience *models.Experience) error {
	var experience models.Experience
	if err := db.DB.First(&experience, id).Error; err != nil {
		return err
	}
	if err := db.DB.Model(&experience).Updates(updatedExperience).Error; err != nil {
		return err
	}
	return nil
}

func DeleteExperience(id uint) error {
	if err := db.DB.Delete(&models.Experience{}, id).Error; err != nil {
		return err
	}
	return nil
}