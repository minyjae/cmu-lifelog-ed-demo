package config

import (
	"fmt"
	"log"
	"os"

	"github.com/minyjae/cmu-life-long-ed-api/internal/adapters/presistence/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase(config *Config) *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort, config.DBSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")

	if shouldRunMigration() {
		runMigration(db)
		if err := CreateSeeds(db, config); err != nil {
			log.Printf("Seeding failed: %v", err)
		}
	} else {
		autoMigrate := os.Getenv("AUTO_MIGRATE")
		appEnv := os.Getenv("APP_ENV")

		if autoMigrate == "false" {
			log.Printf("Skipping database migration (AUTO_MIGRATE=false)")
		} else if appEnv == "production" && autoMigrate != "true" {
			log.Printf("Skipping database migration (production environment, set AUTO_MIGRATE=true to enable)")
		} else {
			log.Printf("Skipping database migration (set AUTO_MIGRATE=true to enable)")
		}

		if err := CreateSeeds(db, config); err != nil {
			log.Printf("Seeding failed: %v", err)
		}
	}

	return db
}

func shouldRunMigration() bool {
	if os.Getenv("AUTO_MIGRATE") == "false" {
		return false
	}

	if os.Getenv("AUTO_MIGRATE") == "true" {
		return true
	}

	if os.Getenv("APP_ENV") == "development" {
		return true
	}

	return false
}

func runMigration(db *gorm.DB) {
	log.Println("Starting database migration...")

	err := db.AutoMigrate(&models.StaffStatus{}, &models.ListQueue{}, &models.Order{}, &models.OrderMapping{}, &models.Users{}, &models.Faculty{}, &models.CourseStatus{}, &models.Role{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migrated successfully")
}

func RunMigrationManual(config *Config) error {
	db := SetupDatabase(config)

	log.Println("Running manual migration...")

	err := db.AutoMigrate(&models.StaffStatus{}, &models.ListQueue{}, &models.Order{}, &models.OrderMapping{}, &models.Users{}, &models.Faculty{}, &models.CourseStatus{}, &models.Role{})
	if err != nil {
		return fmt.Errorf("migration failed: %v", err)
	}

	log.Println("Manual migration completed successfully")
	return nil
}
