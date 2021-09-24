package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"pesthub/adapters/memdb"
	"pesthub/adapters/testmsgs"
	"pesthub/usecases/disorder"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go profile(wg)
	go usecase(wg)

	wg.Wait()
}

func usecase(wg sync.WaitGroup) {
	defer wg.Done()
	store := memdb.NewMemoryDisorderStore()
	messages := testmsgs.NewTestableMessages()
	usecase := disorder.NewRegisterDisorder(store, messages)
	usecase.Execute(&disorder.RegisterDisorderInput{
		Name: "testing",
	})
	fmt.Println("finished usecase")
}

func profile(wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("started listener")
	log.Println(http.ListenAndServe("localhost:6060", nil))
}
