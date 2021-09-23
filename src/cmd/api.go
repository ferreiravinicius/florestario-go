package main

import (
	"pesthub/adapters/api"
	"pesthub/adapters/api/env"
	"pesthub/adapters/memdb"
	"pesthub/adapters/testmsgs"
)

func main() {
	deps := &env.ApiDependencies{
		DisorderStore: memdb.NewMemoryDisorderStore(),
		Messages:      testmsgs.NewTestableMessages(),
	}
	app := api.NewApi(deps)
	app.Listen(":8080")
}
