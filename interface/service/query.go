package service

import (
	"strconv"
	"strings"
)

func reloadDeckQuery(baseURL string, decks int, cardsPerDeck int) string {
	query := strings.Builder{}

	query.WriteString(baseURL)
	query.WriteString("/deck/new/draw/?deck_count=")
	query.WriteString(strconv.Itoa(decks))
	query.WriteString("&count=")
	query.WriteString(strconv.Itoa(decks * cardsPerDeck))

	return query.String()
}
