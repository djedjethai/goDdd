package main

import (
	"fmt"
	"github.com/djedjethai/goDdd/pkg/adding"
	"github.com/djedjethai/goDdd/pkg/http/rest"
	"github.com/djedjethai/goDdd/pkg/listing"
	"github.com/djedjethai/goDdd/pkg/storage/json"
	"github.com/djedjethai/goDdd/pkg/storage/memory"
	"log"
)

type Type int

const (
	JSON Type = iota

	Memory
)

func main() {
	// set up the storage
	storageType := JSON

	var adder adding.Service
	var lister listing.Service

	switch storageType {
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
