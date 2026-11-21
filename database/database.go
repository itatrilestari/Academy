package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	fmt.Println("Database connection successfully opened")
}

func AutoMigrate(models ...interface{}) error {
	log.Println("Starting database migration...")

	for _, model := range models {
		modelName := fmt.Sprintf("%T", model)
		log.Printf("Migrating model: %s", modelName)

		if err := DB.AutoMigrate(model); err != nil {
			log.Printf("Failed to migrate model %s: %v", modelName, err)
			return fmt.Errorf("failed to migrate model %s: %w", modelName, err)
		}

		log.Printf("Successfully migrated model: %s", modelName)
	}

	log.Println("Database migration completed successfully")
	return nil
}
