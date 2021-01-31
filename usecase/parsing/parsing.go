package parsing

import (
	"errors"

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
func RecordToCard(record []string) (model.Card, error) {
	if len(record) != 3 {
		return model.Card{}, errors.New("invalid record")
	}
	return model.Card{
		Value: record[0],
		Suit:  record[1],
		Code:  record[2],
	}, nil
}
