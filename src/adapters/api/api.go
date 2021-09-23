package api

import (
	"pesthub/adapters/api/env"
	"pesthub/adapters/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func NewApi(deps *env.ApiDependencies) *fiber.App {
	setupDependencies(deps)

	config := fiber.Config{
		AppName:      "Florestario",
		ErrorHandler: ErrorHandler,
	}
	app := fiber.New(config)
	setupHandlers(app)
	return app
}

func setupHandlers(app *fiber.App) {
	app.Post("/disorders", handlers.RegisterDisorderHandler)
}

func setupDependencies(deps *env.ApiDependencies) {
	// todo: check all deps and panics if not exists
	env.Deps = deps
}
