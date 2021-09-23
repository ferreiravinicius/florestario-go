package main

import "pesthub/adapters/fiberapi"

func main() {

	app := fiberapi.NewApi()
	app.Listen(":8080")

	// store := memdb.NewMemoryDisorderStore()
	// messages := testmsgs.NewTestableMessages()

	// server := api.NewApi(&config.ApiEnv{
	// 	DisorderStore: store,
	// 	Messages:      messages,
	// })
}
