package fiberapi

import (
	"log"
	"net/http"
	"pesthub/failures"

	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Cause   string `json:"cause,omitempty"`
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	switch err := err.(type) {
	case failures.UseCaseError:
		return ctx.Status(http.StatusUnprocessableEntity).JSON(ErrorResponse{Message: err.Message})
	case failures.InternalError:
		log.Println("Internal error: ", err)
		return ctx.Status(http.StatusInternalServerError).JSON(ErrorResponse{Message: "Oops! Internal error!"})
	default:
		log.Println("Undefined error: ", err)
		return ctx.Status(http.StatusInternalServerError).JSON(ErrorResponse{Message: "Oops! Something went wrong!"})
	}
}
