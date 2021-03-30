package adding

import (
	"fmt"

	"errors"
	"github.com/djedjethai/goDdd/pkg/listing"
)

var errDuplicate = errors.New("beer already exist")

// Service provide beer adding operation
type Service interface {
	AddBeer(...Beer) error
	AddSampleBeers([]Beer)
}

// Repository provide access to beer repository
type Repository interface {
	// add a given beer to the repository
	AddBeer(Beer) error
	// get all beers from the repository
	GetAllBeers() []listing.Beer
}

type service struct {
	r Repository
}

// NewService create an adding service with all the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddBeer persiste in the given beer storage
func (s *service) AddBeer(b ...Beer) error {
	// make sure we don't duplicate
	existingBeers := s.r.GetAllBeers()
	for _, bb := range b {
		for _, e := range existingBeers {
			if bb.Abv == e.Abv &&
				bb.Brewery == e.Brewery &&
				bb.Name == e.Name {
				return errDuplicate
			}
		}
	}

	for _, beer := range b {
		err := s.r.AddBeer(beer) // i add my own err :)
		if err != nil {
			fmt.Println("err when adding beer to the repository")
		}
	}

	return nil
}

func (s *service) AddSampleBeers(b []Beer) {
	for _, bb := range b {
		err := s.r.AddBeer(bb)
		if err != nil {
			fmt.Println("err when adding the sampleBeers to the repos")
		}
	}
}
