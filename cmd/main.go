package main

import (
	"go-fiber-jwt/database"
	"go-fiber-jwt/models"
	"go-fiber-jwt/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// @title API
// @version 1.0
// @description This is the API documentation
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email your-email@domain.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
// @schemes http
func main() {
	database.ConnectDB()

	// Run database migrations
	if err := database.AutoMigrate(
		&models.User{},
		&models.MasterEvalCategory{},
		&models.MasterEvalQuestion{},
		&models.MasterEvalAnswer{},
		&models.EvaluationAnswer{},
		&models.EvaluationCourse{},
	); err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024, // 100 MB
	})

	app.Use(logger.New(logger.Config{
		Format: "${time} | ${status} | ${latency} | ${ip} | ${method} ${path} | err=${error}\n",
	}))
	app.Use(cors.New())

	// Swagger route
	app.Get("/swagger/*", swagger.New(swagger.Config{
		Title:                    "API Documentation",
		DeepLinking:              false,
		DocExpansion:             "none",
		DefaultModelsExpandDepth: -1,
	}))

	routes.SetupRoutes(app)

	app.Listen(":8000")
}
