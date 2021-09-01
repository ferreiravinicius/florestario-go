package main

import (
	"fmt"
	"pesthub/adapters/database"
	"pesthub/features/pests"
)

func main() {
	depsCreate := pests.CreateDeps{
		InsertCommand: database.Insert,
	}

	id, err := pests.Create(depsCreate, &pests.CreateInput{CommonName: "OurName"})
	if err != nil {
		panic(err)
	}

	fmt.Println("Returned id: %v", id)
}
