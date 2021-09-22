package main

import (
	"pesthub/adapters/api"
	"pesthub/adapters/api/config"
	"pesthub/adapters/memdb"
	"pesthub/adapters/testmsgs"
)

func main() {

	store := memdb.NewMemoryDisorderStore()
	messages := testmsgs.NewTestableMessages()

	server := api.NewApi(&config.ApiEnv{
		DisorderStore: store,
		Messages:      messages,
	})

	server.Run()
}
