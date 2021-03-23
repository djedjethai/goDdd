package json

import (
	"encoding/json"
	"fmt"
	"github.com/djedjethai/goDdd/pkg/storage"
	"log"
	"path"
	"runtime"
	"time"

	"github.com/djedjethai/goDdd/pkg/adding"
	"github.com/djedjethai/goDdd/pkg/listing"
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

func (s *Storage) GetBeer(id string) (listing.beer, error) {
	var b Beer
	var beer listing.Beer

	if err := s.db.Read(CollectionBeer, id, &b); err != nil {
		return beer, listing.ErrNotFound
	}

	beer.ID = b.ID
	beer.Name = b.Name
	beer.Brewery = b.Brewery
	beer.Abv = b.Abv
	beer.ShortDesc = b.ShortDesc
	beer.Created = b.Created

	return beer, nil
}

func (s *storage) GetAllBeers() []listing.Beer {
	list := []listing.Beer{}

	records, err := s.db.ReadAll(CollectionBeer)
	if err != nil {
		return list
	}

	for _, r := range records {
		var b Beer
		var beer listing.Beer

		if err := json.Unmarshal([]Byte(r), &b); err != nil {
			return list
		}

		beer.ID = b.ID
		beer.Name = b.Name
		beer.Brewery = b.Brewery
		beer.Abv = b.Abv
		beer.ShortDesc = b.ShortDesc
		beer.Created = b.Created

		list = append(list, beer)
	}

	return list
}
