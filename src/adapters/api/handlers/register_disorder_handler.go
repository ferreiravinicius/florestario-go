package handlers

import (
	"net/http"
	"pesthub/usecases/disorder"

	"github.com/gofiber/fiber/v2"
)

func RegisterDisorderHandler(ctx *fiber.Ctx) error {
	var data disorder.RegisterDisorderInput
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	output, err := disorder.RegisterDisorder(&data)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(&output)
}
