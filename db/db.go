package db

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"backend/config"
	"backend/models"
)

var DB *gorm.DB

func InitDB() {
	dbType := config.GetEnv("DB_TYPE", "sqlite")
	var err error

	if dbType == "postgres" {
		host := config.GetEnv("POSTGRES_HOST", "localhost")
		user := config.GetEnv("POSTGRES_USER", "postgres")
		password := config.GetEnv("POSTGRES_PASSWORD", "")
		dbName := config.GetEnv("POSTGRES_DB", "portfolio")
		port := config.GetEnv("POSTGRES_PORT", "5432")
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		DB, err = gorm.Open(sqlite.Open("portfolio.db"), &gorm.Config{})
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection initialized")
}

func AutoMigrate() {
	// Check if the admin table exists
	if DB.Migrator().HasTable(&models.Admin{}) {
		// Check if the is_super_admin column exists, then drop it
		if DB.Migrator().HasColumn(&models.Admin{}, "is_super_admin") {
			if err := DB.Migrator().DropColumn(&models.Admin{}, "is_super_admin"); err != nil {
				log.Printf("Error dropping is_super_admin column: %v", err)
			}
		}
	}

	// Perform auto-migration for all models
	if err := DB.AutoMigrate(
		&models.Project{},
		&models.Admin{},
		&models.Resume{},
		&models.PortfolioUser{},
		&models.Experience{},
		&models.Education{},
	); err != nil {
		log.Fatalf("Failed to auto-migrate models: %v", err)
	}
}

func SeedDefaultAdmin() {
	// Check if any admin exists
	var count int64
	if err := DB.Model(&models.Admin{}).Count(&count).Error; err != nil {
		log.Printf("Error checking admin count: %v", err)
		return
	}

	if count == 0 {
		 // Extract default admin credentials from environment variables
		defaultUsername := config.GetEnv("DEFAULT_ADMIN_USERNAME", "admin")
		defaultPassword := config.GetEnv("DEFAULT_ADMIN_PASSWORD", "admin123")

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Error hashing default admin password: %v", err)
			return
		}

		defaultAdmin := models.Admin{
			Username: defaultUsername,
			Password: string(hashedPassword),
			Role:     models.RoleSuperSuperAdmin,
		}

		if err := DB.Create(&defaultAdmin).Error; err != nil {
			log.Printf("Error creating default admin: %v", err)
			return
		}

		log.Println("Default admin created successfully")
	} else {
		log.Println("Admin already exists, skipping seeding")
	}
}