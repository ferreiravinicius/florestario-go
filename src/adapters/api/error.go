package api

import (
	"net/http"
	"pesthub/env"
	"pesthub/failures"

	"github.com/gofiber/fiber/v2"
)

var (
	MsgFriendlyInternalError   = env.MessageProvider.Get("friendly.internal.error")
	MsgFriendlyUnexpectedError = env.MessageProvider.Get("friendly.unexpected.error")
)

type ErrorResponse struct {
	Message   string `json:"message"`
	Cause     string `json:"cause,omitempty"`
	FieldName string `json:"field_name,omitempty"`
}

// TODO: log
func ErrorHandler(ctx *fiber.Ctx, err error) error {

	status := http.StatusInternalServerError
	var response ErrorResponse

	switch err := err.(type) {
	case failures.UseCaseError:
		response = ErrorResponse{Message: err.Message}
		status = http.StatusUnprocessableEntity
	case failures.InternalError:
		response = ErrorResponse{Message: MsgFriendlyInternalError}
	case failures.ValidationError:
		response = ErrorResponse{Message: err.Message, FieldName: err.Field}
		status = http.StatusUnprocessableEntity
	default:
		response = ErrorResponse{Message: MsgFriendlyUnexpectedError}
	}

	return ctx.Status(status).JSON(&response)
}
