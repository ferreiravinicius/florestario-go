package restapi

import (
	"fmt"
	"pesthub/adapters/memdb"
	"pesthub/adapters/restapi/handlers"
	"pesthub/adapters/testmsgs"
	"pesthub/env"

	"github.com/gin-gonic/gin"
)

func NewApi() {
	router := gin.Default()

	env.DisorderStore = memdb.NewMemoryDisorderStore()
	env.Messages = testmsgs.NewTestableMessages()

	router.POST("/disorders", H(handlers.RegisterDisorder))
	router.GET("/disorders", FindAllDisordersHandler)

	router.Run(":8080")
}

func FindAllDisordersHandler(c *gin.Context) {
	all, _ := env.DisorderStore.FindAll()
	fmt.Println(all)
	c.JSON(200, all)
}
