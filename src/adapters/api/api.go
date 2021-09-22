package api

import (
	"fmt"
	"pesthub/adapters/api/config"
	"pesthub/adapters/api/handlers"

	"github.com/gin-gonic/gin"
)

type Api struct {
	router *gin.Engine
}

func NewApi(env *config.ApiEnv) *Api {

	config.SetEnv(env)

	router := gin.Default()
	return &Api{
		router,
	}
}

func (api *Api) Run() {
	api.registerHandlers()
	api.router.Run(":8080")
}

func (api *Api) registerHandlers() {
	api.router.POST("/disorders", ApiHandler(handlers.RegisterDisorderHandler).Gin())
	api.router.GET("/disorders", FindAllDisordersHandler)
}

func FindAllDisordersHandler(c *gin.Context) {
	env := config.Env()
	all, _ := env.DisorderStore.FindAll()
	fmt.Println(all)
	c.JSON(200, all)
}
