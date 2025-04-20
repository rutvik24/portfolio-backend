package services

import (
	"backend/db"
	"backend/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CreateAdmin(admin *models.Admin) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return err
	}
	admin.Password = string(hashedPassword)
	if err := db.DB.Create(admin).Error; err != nil {
		log.Printf("Error creating admin: %v", err)
		return err
	}
	return nil
}

func AuthenticateAdmin(username, password string) (*models.Admin, error) {
	var admin models.Admin
	if err := db.DB.Where("username = ?", username).First(&admin).Error; err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return nil, err
	}
	return &admin, nil
}

func GetAdminByID(id uint) (*models.Admin, error) {
	var admin models.Admin
	if err := db.DB.First(&admin, id).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func DeleteAdmin(id uint) error {
	if err := db.DB.Delete(&models.Admin{}, id).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAdminSession(adminID uint, sessionToken string) error {
	if err := db.DB.Model(&models.Admin{}).Where("id = ?", adminID).Update("session_token", sessionToken).Error; err != nil {
		return err
	}
	return nil
}