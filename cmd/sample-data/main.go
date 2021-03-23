package main

import (
	"fmt"

	"github.com/djedjethai/goDdd/cmd/adding"
	"github.com/djedjethai/goDdd/cmd/storage/json"
)

func main() {

	var adder adding.Service

	// err handling omitted for simplicity
	s, _ := json.NewStorage()

	adder = adding.NewService(s)

	// add the sample data
	adder.AddSampleBeers(DefaultBeers)

	fmt.Println("Finished adding sample Beer datas")
}
