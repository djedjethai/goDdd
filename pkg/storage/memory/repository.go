package memory

import (
	// "fmt"
	"github.com/djedjethai/goDdd/pkg/storage"
	"log"
	"time"

	"github.com/djedjethai/goDdd/pkg/adding"
	"github.com/djedjethai/goDdd/pkg/listing"
)

// Memory storage keep data in memory
type Storage struct {
	beers []Beer
}

func (m *Storage) AddBeer(b adding.Beer) error {
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
	m.beers = append(m.beers, newB)

	return nil
}

func (m *Storage) GetBeer(id string) (listing.Beer, error) {
	var beer listing.Beer

	for i := range m.beers {
		if m.beers[i].ID == id {
			beer.ID = id
			beer.Name = m.beers[i].Name
			beer.Brewery = m.beers[i].Brewery
			beer.Abv = m.beers[i].Abv
			beer.ShortDesc = m.beers[i].ShortDesc
			beer.Created = m.beers[i].Created

			return beer, nil
		}
	}

	return beer, listing.ErrNotFound
}

func (m *Storage) GetAllBeers() []listing.Beer {
	var beers []listing.Beer

	for i := range m.beers {
		beer := listing.Beer{
			ID:        m.beers[i].ID,
			Name:      m.beers[i].Name,
			Brewery:   m.beers[i].Brewery,
			Abv:       m.beers[i].Abv,
			ShortDesc: m.beers[i].ShortDesc,
			Created:   m.beers[i].Created,
		}

		beers = append(beers, beer)
	}

	return beers
}
