package rest

import (
	"fmt"
	"pesthub/adapters/memdb"
	"pesthub/adapters/testmsgs"
	"pesthub/env"
	"pesthub/usecases/disorder"

	"github.com/gin-gonic/gin"
)

type Api struct {
}

func main() {
	NewApi()
}

func NewApi() {
	router := gin.Default()

	env.DisorderStore = memdb.NewMemoryDisorderStore()
	env.Messages = testmsgs.NewTestableMessages()

	router.POST("/disorders", RegisterDisorderHandler)
	router.GET("/disorders", FindAllDisordersHandler)

	router.Run(":8080")
}

func FindAllDisordersHandler(c *gin.Context) {
	all, _ := env.DisorderStore.FindAll()
	fmt.Println(all)
	c.JSON(200, all)
}

func RegisterDisorderHandler(c *gin.Context) {

	type input struct {
		Name string `json:"name"`
	}

	var data input
	_ = c.BindJSON(&data)

	output, err := disorder.RegisterDisorder(&disorder.RegisterDisorderInput{
		Name: data.Name,
	})

	if err != nil {
		c.JSON(401, gin.H{"message": "something went wrong", "error": err})
		return
	}

	c.JSON(201, gin.H{"code": output.Code})

}
