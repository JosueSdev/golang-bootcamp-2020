package rules

type table struct {
	amountOfDecks int
	cardsPerDeck  int
}

//Table describes the configuration of the game table
type Table interface {
	AmountOfDecks() int
	CardsPerDeck() int
}

//BlackjackTable constains blackjack's setup rules
var BlackjackTable Table = table{6, 52}

//AmountOfDecks returns how many decks are being used in the table
func (t table) AmountOfDecks() int {
	return t.amountOfDecks
}

//CardsPerDeck returns the size of a single deck
func (t table) CardsPerDeck() int {
	return t.cardsPerDeck
}
