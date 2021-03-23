package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	// "github.com/djedjethai/goDdd/pkg/adding"
	"github.com/djedjethai/goDdd/pkg/listing"
)

func Handler(a adding.Service, l listing.Service) http.Handler {
	router := httprouter.New()

	router.GET("/beers", getBeers(l))
	router.GET("/beers/:id", getBeer(l))
}

func getBeers(s listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list := s.GetBeers()
		json.NewEncoder(w).Encode(list)
	}
}

func getBeer(s listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		beer, err := s.GetBeer(p.ByName("id"))
		if err == listing.ErrNotFound {
			http.Error(w, "the beer you requested do not exist.", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(beer)
	}
}
