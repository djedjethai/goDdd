package memory

import (
	"fmt"
	"github.com/djedjethai/ddd/pkg/storage"
	"log"
	"time"

	"github.com/djedjethai/ddd/pkg/adding"
	"github.com/djedjethai/ddd/pkg/listing"
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

	newB := Beer {
		ID:
	}
}
