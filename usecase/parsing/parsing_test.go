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
	card := RecordToCard(AceRecord)

	if card.Code != AceCard.Code {
		t.Fail()
	}
	if card.Suit != AceCard.Suit {
		t.Fail()
	}
	if card.Value != AceCard.Value {
		t.Fail()
	}
}
