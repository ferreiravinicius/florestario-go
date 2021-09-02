package main

import "pesthub/features/digitalaccount"

func main() {

	e := digitalaccount.NewCreateOffer(nil, nil)
	e.Execute(nil)

	// depsCreate := pest.CreateDeps{
	// 	InsertCommand: database.Insert,
	// }

	// id, err := pest.Create(depsCreate, &pest.CreateInput{CommonName: "OurName"})
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Returned id: ", id)
}
