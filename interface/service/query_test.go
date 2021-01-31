package service

import (
	"fmt"
	"testing"

	"github.com/JosueSdev/golang-bootcamp-2020/domain/rules"
)

const BlackjackDeckReloadQuery = "https://deckofcardsapi.com/api/deck/new/draw/?deck_count=6&count=312"

func TestRealoadDeckQuery(t *testing.T) {
	query := reloadDeckQuery(
		"https://deckofcardsapi.com/api",
		rules.BlackjackTable.AmountOfDecks(),
		rules.BlackjackTable.CardsPerDeck(),
	)
	if query != BlackjackDeckReloadQuery {
		fmt.Println(query)
		t.Fail()
	}
}
