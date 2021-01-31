package rules

import "testing"

var testTable Table = table{
	amountOfDecks: 2,
	cardsPerDeck:  52,
}

func TestAmountOfDecks(t *testing.T) {
	if testTable.AmountOfDecks() != 2 {
		t.Fail()
	}
}

func TestCardsPerDeck(t *testing.T) {
	if testTable.CardsPerDeck() != 52 {
		t.Fail()
	}
}
