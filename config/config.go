package config

import "time"

type api struct {
	BaseURL string
	Timeout time.Duration
}

type server struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

//DeckFileName holds the name of the csv deck file
const DeckFileName = "deck.csv"

//CardsAPI contains the configuration
var CardsAPI = api{
	BaseURL: "https://deckofcardsapi.com/api",
	Timeout: 3 * time.Second,
}

//Server contains the server's configuration
var Server = server{
	Addr:         ":8080",
	ReadTimeout:  6 * time.Second,
	WriteTimeout: 6 * time.Second,
}
