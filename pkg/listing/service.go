package listing

import (
	"errors"
)

var ErrNotFound = errors.New("beer not found")

// Repository provide access to the beer and review storage
type Repository interface {
	// get from the repos the matched beer(using id)
	GetBeer(string) (Beer, error)

	// all beer from storage
	GetAllBeers() []Beer
}

// provide beer and listing operations
type Service interface {
	GetBeer(string) (Beer, error)
	GetBeers() []Beer
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetBeers() []Beer {
	return s.r.GetAllBeers()
}

func (s *service) GetBeer(id string) (Beer, error) {
	return s.r.GetBeer(id)

}
