package main

import (
	"pesthub/adapters/api"
	"pesthub/adapters/api/apideps"
	"pesthub/adapters/memdb"
	"pesthub/adapters/testmsgs"
)

func main() {
	apideps.DisorderStore = memdb.NewMemoryDisorderStore()
	apideps.Messages = testmsgs.NewTestableMessages()

	server := api.NewApi()
	server.Run()
}
