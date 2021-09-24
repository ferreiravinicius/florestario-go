package main

import (
	"pesthub/adapters/memdb"
	"pesthub/adapters/testmsgs"
	"pesthub/usecases/disorder"
)

var store = memdb.NewMemoryDisorderStore()
var messages = testmsgs.NewTestableMessages()

func main() {
	usecase := disorder.NewRegisterDisorder(store, messages)
	usecase.Execute(&disorder.RegisterDisorderInput{
		Name: "testing",
	})
}
