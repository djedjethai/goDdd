package json

import (
	"encoding/json"
	"fmt"
	"github.com/djedjethai/ddd/storage"
	"log"
	"path"
	"runtime"
	"time"

	"github.com/djedjethai/ddd/adding"
	"github.com/djedjethai/ddd/listing"
	"github.com/nanobox.io/golang-scribble"
)

const (
	dir = "/data"

	// identifier for the json collection of beer
	CollectionBeer = "beers"
)

// Storage stores beer data in JSON files
type Storage struct {
	db *scribble.Driver
}

// NewStorage return a new JSON Storage
func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	_, filename, _, _ := runtime.Caller(0)
	p := path.Dir(filename)

	s.db, err = scribble.New(p+dir, nil)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// AddBeer saves the given beer to the repository
func (s *Storage) AddBeer(b adding.Beer) error {
	id, err := storage.GetID("beer")
	if err != nil {
		log.Fatal(err)
	}

	newB := Beer{
		ID:        id,
		Created:   time.Now(),
		Name:      b.Name,
		Brewery:   b.Brewery,
		Abv:       b.Abv,
		ShortDesc: b.ShortDesc,
	}

	if err := s.db.Write(CollectionBeer, newB.ID, newB); err != nil {
		return err
	}

	return nil
}
