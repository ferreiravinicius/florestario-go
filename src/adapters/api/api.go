package api

import (
	"fmt"
	"pesthub/adapters/api/apideps"
	"pesthub/adapters/api/handlers"

	"github.com/gin-gonic/gin"
)

type Api struct {
	router *gin.Engine
}

func NewApi() *Api {
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
	api.router.POST("/disorders", H(handlers.RegisterDisorder))
	api.router.GET("/disorders", FindAllDisordersHandler)
}

func FindAllDisordersHandler(c *gin.Context) {
	all, _ := apideps.DisorderStore.FindAll()
	fmt.Println(all)
	c.JSON(200, all)
}

// router.POST("/disorders", H(handlers.RegisterDisorder))
// router.GET("/disorders", FindAllDisordersHandler)
