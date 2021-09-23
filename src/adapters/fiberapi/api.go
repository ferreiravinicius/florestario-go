package fiberapi

import (
	"github.com/gofiber/fiber/v2"
)

func NewApi() *fiber.App {
	config := fiber.Config{
		AppName:      "Florestario",
		ErrorHandler: ErrorHandler,
	}
	app := fiber.New(config)
	setupHandlers(app)
	return app
}

func setupHandlers(app *fiber.App) {
	app.Get("/health")
}
