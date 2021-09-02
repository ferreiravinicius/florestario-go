package main

import (
	"pesthub/features/contadigital/enviarlink"
)

func main() {

	enviarlink.EnviarWhatsapp()

	// depsCreate := pest.CreateDeps{
	// 	InsertCommand: database.Insert,
	// }

	// id, err := pest.Create(depsCreate, &pest.CreateInput{CommonName: "OurName"})
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Returned id: ", id)
}
