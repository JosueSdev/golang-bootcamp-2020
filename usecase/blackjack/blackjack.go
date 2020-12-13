package blackjack

import (
	"strconv"

	"github.com/JosueSdev/golang-bootcamp-2020/domain/model"
	"github.com/JosueSdev/golang-bootcamp-2020/domain/rules"
)

//CardToValues extracts the possible values of a card. Returns an error if an inavlid card is gives as input.
func CardToValues(card model.Card) ([2]int, error) {
	switch {
	case card.Value == "ACE":
		return [2]int{1, 10}, nil
	case card.Value == "KING", card.Value == "QUEEN", card.Value == "KING":
		return [2]int{10, 0}, nil
	default:
		val, err := strconv.Atoi(card.Value)
		return [2]int{val, 0}, err
	}
}

//CalculateScore returns the best possible score for a set of Cards. Returns an error if an inavlid card is gives as input.
func CalculateScore(hand []model.Card) (score int, gameStatus string, err error) {
	aces := 0

	for _, card := range hand {
		values, err := CardToValues(card)
		if err != nil {
			return 0, "", err
		}
		score += values[0]
		if values[1] > 0 {
			aces++
		}
	}

	isBlackjack, isFoul := checkScore(score)
	if isBlackjack {
		return score, rules.EnglishStatus.Win(), nil
	}
	if isFoul {
		return score, rules.EnglishStatus.Foul(), nil
	}

	for i := 0; i < aces; i++ {
		isBlackjack, isFoul := checkScore(score + 10)
		if isBlackjack {
			return 21, rules.EnglishStatus.Win(), nil
		}
		if isFoul {
			break
		}
		score += 10
	}

	return score, rules.EnglishStatus.Ok(), nil
}

//IsPositionInTableBounds validates that a given card position can exist given the amount of cards in the table
func IsPositionInTableBounds(i int) bool {
	switch {
	case i >= rules.BlackjackTable.AmountOfDecks()*rules.BlackjackTable.CardsPerDeck():
		return false
	case i < 0:
		return false
	default:
		return true
	}
}

func checkScore(score int) (isBlackjack bool, isFoul bool) {
	if score == 21 {
		return true, false
	}
	if score > 21 {
		return false, true
	}
	return false, false
}