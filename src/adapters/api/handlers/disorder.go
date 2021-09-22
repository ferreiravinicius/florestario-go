package handlers

import (
	"pesthub/adapters/api/apideps"
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

	usecase := disorder.NewRegisterDisorder(
		apideps.DisorderStore,
		apideps.Messages,
	)

	output, err := usecase.Execute(&disorder.RegisterDisorderInput{
		Name: data.Name,
	})
	if err != nil {
		return err
	}

	resp := &response{Code: output.Code}
	ctx.JSON(201, resp)

	return nil
}
