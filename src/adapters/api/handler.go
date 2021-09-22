package api

import (
	"log"
	"net/http"
	"pesthub/failures"

	"github.com/gin-gonic/gin"
)

// Our handler func to better error handling
type OurHandlerFunc func(ctx *gin.Context) error

// Adapts OurHandlerFunc to gin.HandlerFunc
func H(handler OurHandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := handler(ctx); err != nil {
			HandleError(ctx, err)
		}
	}
}

// Handle errors in one place
func HandleError(ctx *gin.Context, err error) {
	switch err := err.(type) {
	case failures.UseCaseError:
		ctx.JSON(http.StatusUnprocessableEntity, ErrorResponse{
			Message: err.Message,
		})
	case failures.InternalError:
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: "Oops!",
		})
		log.Println("[?] Internal error: ", err)
	default:
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: "Oops!!!",
		})
		log.Println("[!] Unknown error: ", err)
	}
}

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}
