package main

import (
	"pesthub/adapters/api"
	"pesthub/adapters/memdb"
	"pesthub/adapters/testmsgs"
	"pesthub/env"
)

func init() {
	env.DisorderStore = memdb.NewMemoryDisorderStore()
	env.MessageProvider = testmsgs.NewTestableMessageProvider()
}

func main() {
	app := api.NewApi()
	app.Listen(":8080")
}
