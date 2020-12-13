package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JosueSdev/golang-bootcamp-2020/config"
	"github.com/JosueSdev/golang-bootcamp-2020/infrastructure/client"
	"github.com/JosueSdev/golang-bootcamp-2020/infrastructure/datastore"
	"github.com/JosueSdev/golang-bootcamp-2020/infrastructure/router"
	"github.com/JosueSdev/golang-bootcamp-2020/interface/app"
)

func main() {
	deckCSV, err := datastore.NewCSVDeck()

	defer func() {
		err = deckCSV.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatal(err)
	}

	deckClient := client.NewCardsClient()

	app := app.NewApp(deckClient, deckCSV)
	router := router.Build(app)

	s := &http.Server{
		Addr:         config.Server.Addr,
		Handler:      router,
		ReadTimeout:  config.Server.ReadTimeout,
		WriteTimeout: config.Server.WriteTimeout,
	}
	fmt.Printf("listening at %s\n", config.Server.Addr)
	log.Fatal(s.ListenAndServe())
}
