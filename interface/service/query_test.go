package service

import (
	"testing"

	"github.com/JosueSdev/golang-bootcamp-2020/config"
	"github.com/JosueSdev/golang-bootcamp-2020/domain/rules"
)

const BlackjackDeckReloadQuery = "https://deckofcardsapi.com/api/deck/new/draw/?deck_count=6&count=312"

func TestRealoadDeckQuery(t *testing.T) {
	query := reloadDeckQuery(
		config.CardsAPI.BaseURL,
		rules.BlackjackTable.AmountOfDecks(),
		rules.BlackjackTable.CardsPerDeck(),
	)
	if query != BlackjackDeckReloadQuery {
		t.Fail()
	}
}
