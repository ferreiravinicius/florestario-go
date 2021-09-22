package handlers

import (
	"pesthub/adapters/api/config"
	"pesthub/usecases/disorder"

	"github.com/gin-gonic/gin"
)

type request struct {
	Name string `json:"name"`
}

type response struct {
	Code string `json:"code"`
}

func RegisterDisorderHandler(ctx *gin.Context) error {

	var data request
	ctx.BindJSON(&data)

	env := config.Env()
	usecase := disorder.NewRegisterDisorder(
		env.DisorderStore,
		env.Messages,
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
