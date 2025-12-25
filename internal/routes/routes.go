package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-backend/internal/handler"
	"go-backend/internal/middleware"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	// Middleware
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger())

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// User routes
	api := app.Group("/users")
	api.Post("", userHandler.CreateUser)
	api.Get("", userHandler.ListUsers) // Must come before /:id route
	api.Get("/:id", userHandler.GetUserByID)
	api.Put("/:id", userHandler.UpdateUser)
	api.Delete("/:id", userHandler.DeleteUser)
}
