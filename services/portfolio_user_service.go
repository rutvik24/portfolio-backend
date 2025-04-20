package services

import (
	"backend/db"
	"backend/models"
)

func GetPortfolioUser() (*models.PortfolioUser, error) {
	var user models.PortfolioUser
	if err := db.DB.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdatePortfolioUser(updatedUser *models.PortfolioUser) error {
	var user models.PortfolioUser
	if err := db.DB.First(&user).Error; err != nil {
		return err
	}
	if err := db.DB.Model(&user).Updates(updatedUser).Error; err != nil {
		return err
	}
	return nil
}