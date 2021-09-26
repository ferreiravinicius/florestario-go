package main

import (
	"fmt"
	"pesthub/adapters/api"
	"pesthub/adapters/memdb"
	"pesthub/adapters/testmsgs"
	"pesthub/entities"
	"pesthub/env"
)

func init() {
	env.DisorderStore = memdb.NewMemoryDisorderStore()
	env.MessageProvider = testmsgs.NewTestableMessageProvider()
}

func main() {
	_ = api.NewApi()
	// app.Listen(":8080")

	d := entities.Disorder{}
	fmt.Printf("%v \n", d)
	fmt.Println(d.Causer.Name)

}
