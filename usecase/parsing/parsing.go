package parsing

import (
	"github.com/JosueSdev/golang-bootcamp-2020/domain/model"
	"github.com/JosueSdev/golang-bootcamp-2020/domain/rules"
)

//CardsToRecords convert cards into csv writeable records
func CardsToRecords(cards []model.Card) [][]string {
	records := [][]string{
		{
			rules.CardFieldNames.Value,
			rules.CardFieldNames.Suit,
			rules.CardFieldNames.Code,
		},
	}

	for _, card := range cards {
		records = append(records, []string{card.Value, card.Suit, card.Code})
	}

	return records
}

//RecordToCard converts a csv card record into model.Card
func RecordToCard(record []string) model.Card {
	return model.Card{
		Value: record[0],
		Suit:  record[1],
		Code:  record[2],
	}
}

//RecordsToCards converts csv card records into model.Card
func RecordsToCards(records [][]string) []model.Card {
	cards := []model.Card{}

	for _, record := range records {
		cards = append(cards, RecordToCard(record))
	}

	return cards
}
