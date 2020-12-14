package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type api struct {
	BaseURL string
	Timeout time.Duration
}

type server struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var (
	//DeckFileName holds the name of the csv deck file
	DeckFileName string
	//CardsAPI contains the configuration of the external API
	CardsAPI api
	//Server contains the server's configuration
	Server server
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("config file error")
	}

	DeckFileName = viper.GetString("deckfilename")

	CardsAPI = api{
		BaseURL: viper.GetString("cardsapi.baseurl"),
		Timeout: time.Duration(viper.GetInt("cardsapi.timeout")) * time.Second,
	}

	Server = server{
		Addr:         viper.GetString("server.addr"),
		ReadTimeout:  time.Duration(viper.GetInt("server.readtimeout")) * time.Second,
		WriteTimeout: time.Duration(viper.GetInt("server.writetimeout")) * time.Second,
	}
}

//DeckFileName holds the name of the csv deck file
//const DeckFileName = "deck.csv"

//CardsAPI contains the configuration
/*
var CardsAPI = api{
	BaseURL: "https://deckofcardsapi.com/api",
	Timeout: 3 * time.Second,
}
*/

//Server contains the server's configuration
/*
var Server = server{
	Addr:         ":8080",
	ReadTimeout:  6 * time.Second,
	WriteTimeout: 6 * time.Second,
}
*/
