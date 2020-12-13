package app

import (
	"github.com/JosueSdev/golang-bootcamp-2020/infrastructure/client"
	"github.com/JosueSdev/golang-bootcamp-2020/infrastructure/datastore"
	"github.com/JosueSdev/golang-bootcamp-2020/interface/handler"
	"github.com/JosueSdev/golang-bootcamp-2020/interface/service"
)

type app struct {
	blackjackHandler handler.Blackjack
}

//App defines getters for handlers for the supported funtionality
type App interface {
	BlackjackHandler() handler.Blackjack
}

//NewApp constructs an App
func NewApp(cl client.Deck, csv datastore.CSVDeck) App {
	app := new(app)

	app.blackjackHandler = handler.NewBlackjackHandler(service.NewDeckService(cl, csv))

	return app
}

func (a *app) BlackjackHandler() handler.Blackjack {
	return a.blackjackHandler
}
