package router

import (
	"github.com/JosueSdev/golang-bootcamp-2020/interface/handler"
	"github.com/go-chi/chi"
)

type blackJack struct {
	handler handler.Blackjack
}

//BlackjackBuilder provides a subrouter builder
type BlackjackBuilder interface {
	Build() *chi.Mux
}

//NewBlackjackBuilder constructs a BlackjackBuilder with dependency injection
func NewBlackjackBuilder(bh handler.Blackjack) BlackjackBuilder {
	return &blackJack{bh}
}

//Build builds a subrouter for the blackjack resource
func (br *blackJack) Build() *chi.Mux {
	mux := chi.NewRouter()

	mux.Get("/game", br.handler.GetGame)

	mux.Put("/table", br.handler.PutTable)

	return mux
}
