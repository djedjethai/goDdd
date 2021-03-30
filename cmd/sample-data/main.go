package main

import (
	"fmt"

	"github.com/djedjethai/goDdd/pkg/adding"
	"github.com/djedjethai/goDdd/pkg/storage/json"
)

func main() {

	var adder adding.Service

	// err handling omitted for simplicity
	s, _ := json.NewStorage()

	adder = adding.NewService(s)

	// add the sample data
	fmt.Printf("%v", DefaultBeers)
	fmt.Printf("%v", Moncul)
	adder.AddSampleBeers(DefaultBeers)

	fmt.Println("Finished adding sample Beer datas")
}
