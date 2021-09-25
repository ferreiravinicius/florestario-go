package api

import (
	"net/http"
	"pesthub/env"
	"pesthub/failures"

	"github.com/gofiber/fiber/v2"
)

var (
	MsgFriendlyInternalError   = "friendly.internal.error"
	MsgFriendlyUnexpectedError = "friendly.unexpected.error"
)

type ErrorResponse struct {
	Message   string `json:"message"`
	Cause     error  `json:"cause,omitempty"`
	FieldName string `json:"field_name,omitempty"`
}

// TODO: add log, remove cause and manage messages properly
func ErrorHandler(ctx *fiber.Ctx, err error) error {

	code := http.StatusInternalServerError
	var response ErrorResponse

	switch err := err.(type) {
	case failures.UseCaseError:
		response = ErrorResponse{Message: err.Message}
		code = http.StatusUnprocessableEntity
	case failures.ValidationError:
		response = ErrorResponse{Message: err.Message, FieldName: err.Field}
		code = http.StatusUnprocessableEntity
	case failures.InternalError:
		msg := env.MessageProvider.Get(MsgFriendlyInternalError)
		response = ErrorResponse{Message: msg}
	default:
		msg := env.MessageProvider.Get(MsgFriendlyUnexpectedError)
		response = ErrorResponse{Message: msg, Cause: err}
	}

	return ctx.Status(code).JSON(&response)
}
