package services

import (
	"backend/db"
	"backend/models"
	"log"
)

func GetAllEducations() []models.Education {
	var educations []models.Education
	if err := db.DB.Find(&educations).Error; err != nil {
		log.Printf("Error retrieving educations: %v", err)
	}
	return educations
}

func GetEducationByID(id uint) (*models.Education, error) {
	var education models.Education
	if err := db.DB.First(&education, id).Error; err != nil {
		return nil, err
	}
	return &education, nil
}

func CreateEducation(education *models.Education) error {
	if err := db.DB.Create(education).Error; err != nil {
		log.Printf("Error creating education in database: %v", err)
		return err
	}
	return nil
}

func UpdateEducation(id uint, updatedEducation *models.Education) error {
	var education models.Education
	if err := db.DB.First(&education, id).Error; err != nil {
		return err
	}
	if err := db.DB.Model(&education).Updates(updatedEducation).Error; err != nil {
		return err
	}
	return nil
}

func DeleteEducation(id uint) error {
	if err := db.DB.Delete(&models.Education{}, id).Error; err != nil {
		return err
	}
	return nil
}