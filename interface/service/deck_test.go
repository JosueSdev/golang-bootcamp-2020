package service

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestReadIndexedCards(t *testing.T) {
	file, err := os.Open("../../testdata/deck.csv")

	if err != nil {
		t.Fatal(err)
	}

	reader := csv.NewReader(file)
	indexes := []int{2, 0}

	cards, err := readIndexedCards(indexes, reader)

	if err != nil {
		t.Error(err)
	}

	if len(cards) != len(indexes) {
		t.Fail()
	}

	if cards[0].Value != "ACE" && cards[1].Value != "Queen" {
		t.Fail()
	}
}
