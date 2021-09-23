package handlers

import (
	"net/http"
	"pesthub/adapters/api/env"
	"pesthub/usecases/disorder"

	"github.com/gofiber/fiber/v2"
)

func RegisterDisorderHandler(ctx *fiber.Ctx) error {
	var data disorder.RegisterDisorderInput
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	usecase := disorder.NewRegisterDisorder(
		env.Deps.DisorderStore,
		env.Deps.Messages,
	)

	output, err := usecase.Execute(&data)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(&output)
}
