package main

import (
	"fmt"
	"pesthub/adapters/database"
	"pesthub/features/pest"
)

func main() {
	depsCreate := pest.CreateDeps{
		InsertCommand: database.Insert,
	}

	id, err := pest.Create(depsCreate, &pest.CreateInput{CommonName: "OurName"})
	if err != nil {
		panic(err)
	}

	fmt.Println("Returned id: %v", id)
}
