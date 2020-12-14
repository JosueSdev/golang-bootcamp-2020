package blackjack

import (
	"testing"

	"github.com/JosueSdev/golang-bootcamp-2020/domain/model"
	"github.com/JosueSdev/golang-bootcamp-2020/domain/rules"
)

var (
	Ace = model.Card{
		Value: "ACE",
		Suit:  "HEARTS",
		Code:  "AH",
	}
	Jack = model.Card{
		Value: "JACK",
		Suit:  "HEARTS",
		Code:  "JH",
	}
	Queen = model.Card{
		Value: "QUEEN",
		Suit:  "HEARTS",
		Code:  "JH",
	}
	King = model.Card{
		Value: "KING",
		Suit:  "HEARTS",
		Code:  "KH",
	}
	Three = model.Card{
		Value: "3",
		Suit:  "HEARTS",
		Code:  "3H",
	}
	Eight = model.Card{
		Value: "8",
		Suit:  "HEARTS",
		Code:  "8H",
	}
	Bad = model.Card{
		Value: "PIKACHU",
		Suit:  "POKEMON",
		Code:  "PP",
	}
	RegularHand   = []model.Card{Eight, Eight}
	WiningHand    = []model.Card{Ace, Three, Eight}
	FoulHand      = []model.Card{Eight, Eight, Eight}
	DoubleAceHand = []model.Card{Ace, Ace, Eight}
	BadHand       = []model.Card{Ace, Bad}
)

func TestCardToValues(t *testing.T) {
	tests := map[string]struct {
		Card   model.Card
		Values [2]int
		Err    bool
	}{
		"calculates value of simple card": {
			Three,
			[2]int{3, 0},
			false,
		},
		"calculates value of ace": {
			Ace,
			[2]int{1, 10},
			false,
		},
		"calculates value of jack": {
			Jack,
			[2]int{10, 0},
			false,
		},
		"calculates value of queen": {
			Queen,
			[2]int{10, 0},
			false,
		},
		"calculates value of king": {
			King,
			[2]int{10, 0},
			false,
		},
		"fails on bad card": {
			Bad,
			[2]int{0, 0},
			true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			values, err := CardToValues(test.Card)
			if values != test.Values {
				t.Fail()
			}
			if !test.Err && err != nil {
				t.Fail()
			}
		})
	}
}

func TestCheckScore(t *testing.T) {
	tests := map[string]struct {
		Score       int
		IsBlackjack bool
		IsFoul      bool
	}{
		"21 is blackjack": {
			21,
			true,
			false,
		},
		"22 is foul": {
			22,
			false,
			true,
		},
		"20 is ok": {
			20,
			false,
			false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			isBlackjack, isFoul := checkScore(test.Score)
			if isBlackjack != test.IsBlackjack {
				t.Fail()
			}
			if isFoul != test.IsFoul {
				t.Fail()
			}
		})
	}
}

func TestCalculateScore(t *testing.T) {
	tests := map[string]struct {
		Hand   []model.Card
		Score  int
		Status string
		Err    bool
	}{
		"regular hand ok": {
			RegularHand,
			16,
			rules.EnglishStatus.Ok(),
			false,
		},
		"winning hand win": {
			WiningHand,
			21,
			rules.EnglishStatus.Win(),
			false,
		},
		"foul hand foul": {
			FoulHand,
			24,
			rules.EnglishStatus.Foul(),
			false,
		},
		"double ace hand ok": {
			DoubleAceHand,
			19,
			rules.EnglishStatus.Ok(),
			false,
		},
		"bad hand error": {
			BadHand,
			0,
			"",
			true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			score, gameStatus, err := CalculateScore(test.Hand)
			if score != test.Score {
				t.Fail()
			}
			if gameStatus != test.Status {
				t.Fail()
			}
			if !test.Err && err != nil {
				t.Fail()
			}
		})
	}
}
