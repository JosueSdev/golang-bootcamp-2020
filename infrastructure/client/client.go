package client

import (
	"net/http"

	"github.com/JosueSdev/golang-bootcamp-2020/config"
)

type deck struct {
	client  *http.Client
	baseURL string
}

//Deck provides a client configured to consume the Deck of Cards API
type Deck interface {
	Client() *http.Client
	BaseURL() string
}

//NewCardsClient returns an http client configured to make requests to the Deck of Cards API
func NewCardsClient() Deck {
	return &deck{
		&http.Client{Timeout: config.CardsAPI.Timeout},
		config.CardsAPI.BaseURL,
	}
}

func (d *deck) Client() *http.Client {
	return d.client
}

func (d *deck) BaseURL() string {
	return d.baseURL
}
