package fetching

import (
	beerscli "github.com/LuisCusihuaman/golang-introduction/internal"
	"github.com/pkg/errors"
)

type Service interface {
	FetchBeers() ([]beerscli.Beer, error)
	FetchByID(id int) (beerscli.Beer, error)
}

type service struct {
	bR beerscli.BeerRepo
}

func NewService(r beerscli.BeerRepo) Service {
	return &service{r}
}

func (s *service) FetchBeers() ([]beerscli.Beer, error) {
	return s.bR.GetBeers()
}

func (s *service) FetchByID(id int) (beerscli.Beer, error) {
	beers, err := s.FetchBeers()
	if err != nil {
		return beerscli.Beer{}, err
	}
	for _, beer := range beers {
		if beer.ProductID == id {
			return beer, nil
		}
	}
	return beerscli.Beer{}, errors.Errorf("Beer %d not found", id)
}
