package beerscli

import "encoding/json"

type BeerType int

const (
	Unknown BeerType = iota
	Lager
	Malt
	Ale
	FlavouredMalt
	Stout
	Porter
	NonAlcoholic
)

var toString = map[BeerType]string{
	Lager:         "Lager",
	Malt:          "Malt",
	Ale:           "Ale",
	FlavouredMalt: "Flavoured Malt",
	Stout:         "Stout",
	Porter:        "Stout",
	NonAlcoholic:  "Non-Alcoholic",
	Unknown:       "unknown",
}

var toID = map[string]BeerType{
	"Lager":          Lager,
	"Malt":           Malt,
	"Ale":            Ale,
	"Flavoured Malt": FlavouredMalt,
	"Stout":          Stout,
	"Porter":         Porter,
	"Non-Alcoholic":  NonAlcoholic,
	"unknown":        Unknown,
}

func (t *BeerType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*t = toID[j]
	return nil
}

func NewBeerType(t string) *BeerType {
	beerType := toID[t]
	return &beerType
}

type Beer struct {
	ProductID int       `json:"product_id"`
	Name      string    `json:"name"`
	Price     string    `json:"price"`
	BeerID    int       `json:"beer_id"`
	Category  string    `json:"category"`
	Type      *BeerType `json:"type"`
	Brewer    string    `json:"brewer"`
	Country   string    `json:"country"`
}

func NewBeer(productID int, name, category, brewer, country, price string, beerType *BeerType) (b Beer) {
	b = Beer{
		ProductID: productID,
		Name:      name,
		Category:  category,
		Type:      beerType,
		Brewer:    brewer,
		Country:   country,
		Price:     price,
	}
	return
}

type BeerRepo interface {
	GetBeers() ([]Beer, error)
}
