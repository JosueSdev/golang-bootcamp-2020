package parsing

import (
	"reflect"
	"testing"

	"github.com/JosueSdev/golang-bootcamp-2020/domain/model"
	"github.com/JosueSdev/golang-bootcamp-2020/domain/rules"
)

var (
	AceCard = model.Card{
		Value: "ACE",
		Suit:  "HEARTS",
		Code:  "AH",
	}
	AceRecord = []string{
		"ACE",
		"HEARTS",
		"AH",
	}
	EightCard = model.Card{
		Value: "8",
		Suit:  "HEARTS",
		Code:  "8H",
	}
	EightRecord = []string{
		"8",
		"HEARTS",
		"8H",
	}
	BadRecord = []string{
		"trash",
		"taste",
	}
	Cards   = []model.Card{AceCard, EightCard}
	Records = [][]string{
		{
			rules.CardFieldNames.Value,
			rules.CardFieldNames.Suit,
			rules.CardFieldNames.Code,
		},
		AceRecord,
		EightRecord,
	}
)

func TestCardsToRecords(t *testing.T) {
	records := CardsToRecords(Cards)

	//slow
	if !reflect.DeepEqual(records, Records) {
		t.Fail()
	}
}

func TestRecordToCard(t *testing.T) {
	tests := map[string]struct {
		Record []string
		Card   model.Card
		Err    bool
	}{
		"converts ace record": {
			AceRecord,
			AceCard,
			false,
		},
		"fails on shorter record": {
			BadRecord,
			model.Card{},
			true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			card, err := RecordToCard(test.Record)

			if (err != nil) != test.Err {
				t.Fail()
			}
			if card.Code != test.Card.Code {
				t.Fail()
			}
			if card.Suit != test.Card.Suit {
				t.Fail()
			}
			if card.Value != test.Card.Value {
				t.Fail()
			}
		})
	}
}
