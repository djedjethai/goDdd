package main

import (
	"fmt"
	"log"
	"github.com/djedjethai/ddd/pkg/adding"
	"github.com/djedjethai/ddd/pkg/http/rest"
	"github.com/djedjethai/ddd/pkg/listing"
	"github.com/djedjethai/ddd/pkg/storage/json"
	"github.com/djedjethai/ddd/pkg/storage/memory"
)

type Type int 

const (
	JSON.Type = iota

	Memory
)


func main() {
	// set up the storage
	storageType := JSON

	var adder adding.service
	var lister listing.service

	switch storageType{
	case Memory:
		s := new(memory.Storage)

		adder = adding.NewService(s)
		lister = listing.NewService(s)

	case JSON:
		s, _ := json.NewStorage()

		adder = adding.NewService(s)
		lister = listing.NewService(s)
	}

	// set up the http server
	router := rest.handler(adder, lister)

	fmt.Println("server listen on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
