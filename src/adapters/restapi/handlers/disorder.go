package handlers

import (
	"pesthub/usecases/disorder"

	"github.com/gin-gonic/gin"
)

func RegisterDisorder(ctx *gin.Context) error {

	type request struct {
		Name string `json:"name"`
	}

	type response struct {
		Code string `json:"code"`
	}

	var data request
	ctx.BindJSON(&data)

	output, err := disorder.RegisterDisorder(&disorder.RegisterDisorderInput{
		Name: data.Name,
	})
	if err != nil {
		return err
	}

	resp := &response{Code: output.Code}
	ctx.JSON(201, resp)

	return nil
}
