package routes

import (
	"go-fiber-jwt/handlers"
	"go-fiber-jwt/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Public
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "COOFISDEV ACADEMY API",
			"version": "1.0.0",
		})
	})

	auth := app.Group("/api/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	// Protected
	api := app.Group("/api", middleware.AuthMiddleware())

	// Evaluation
	evaluation := api.Group("/evaluation")
	evaluation.Post("/", handlers.CreateEvaluation)
	evaluation.Get("/:id", handlers.GetListEvaluationAnswer)

	// Admin
	admin := api.Group("/admin")

	// Master Evaluation Category
	masterCategory := admin.Group("/master/category")
	masterCategory.Post("/", handlers.CreateMasterEvalCategory)
	masterCategory.Get("/", handlers.GetAllMasterEvalCategory)
	masterCategory.Put("/:id", handlers.UpdateMasterEvalCategory)
	masterCategory.Delete("/:id", handlers.DeleteMasterEvalCategory)

	// Master Evaluation Question
	masterEvalQuestion := admin.Group("/master/evaluation/question")
	masterEvalQuestion.Post("/", handlers.CreateMasterEvalQuestion)
	masterEvalQuestion.Get("/", handlers.MasterEvalQuestion)
	masterEvalQuestion.Put("/:id", handlers.UpdateMasterEvalQuestion)
	masterEvalQuestion.Delete("/:id", handlers.DeleteMasterEvalQuestion)

	// Master Evaluation Answer
	masterEvalAnswer := admin.Group("/master/evaluation/answer")
	masterEvalAnswer.Post("/", handlers.CreateMasterEvalAnswer)
	masterEvalAnswer.Get("/", handlers.GetAllMasterEvalAnswer)
	masterEvalAnswer.Put("/:id", handlers.UpdateMasterEvalAnswer)
	masterEvalAnswer.Delete("/:id", handlers.DeleteMasterEvalAnswer)
}
