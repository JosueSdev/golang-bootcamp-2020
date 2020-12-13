package router

import (
	"github.com/JosueSdev/golang-bootcamp-2020/interface/app"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//Build builds the api's router
func Build(app app.App) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/blackjack", NewBlackjackBuilder(app.BlackjackHandler()).Build())

	return r
}
