package ontario

import (
	"encoding/json"
	"fmt"
	beerscli "github.com/LuisCusihuaman/golang-introduction/internal"
	"github.com/LuisCusihuaman/golang-introduction/internal/errors"
	"io/ioutil"
	"net/http"
)

const (
	productsEndpoint = "/products"
	ontarioURL       = "http://ontariobeerapi.ca"
)

type beerRepo struct {
	url string
}

// NewOntarioRepository fetch beers data from csv
func NewOntarioRepository() beerscli.BeerRepo {
	return &beerRepo{url: ontarioURL}
}

func (b *beerRepo) GetBeers() (beers []beerscli.Beer, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v", b.url, productsEndpoint))
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error getting response to %s", productsEndpoint)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error reading the response from %s", productsEndpoint)
	}

	err = json.Unmarshal(contents, &beers)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "can't parsing response into beers")
	}
	return
}
