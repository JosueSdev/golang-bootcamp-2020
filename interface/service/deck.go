package service

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"sort"

	"github.com/JosueSdev/golang-bootcamp-2020/domain/model"
	"github.com/JosueSdev/golang-bootcamp-2020/domain/rules"
	"github.com/JosueSdev/golang-bootcamp-2020/infrastructure/client"
	"github.com/JosueSdev/golang-bootcamp-2020/infrastructure/datastore"
	"github.com/JosueSdev/golang-bootcamp-2020/usecase/parsing"
)

type deck struct {
	API client.Deck
	CSV datastore.CSVDeck
}

//Deck interfaces with the API and CSV to provide Deck services
type Deck interface {
	ReloadDeck() error
	GetHand([]int) ([]model.Card, error)
}

//NewDeckService constructs a deck service with dependency injection
func NewDeckService(cl client.Deck, csv datastore.CSVDeck) Deck {
	return &deck{cl, csv}
}

//ReloadDeck fetches a new deck an uses it to write/replace the csv storage
func (d *deck) ReloadDeck() error {
	//request a new deck
	resp, err := d.API.Client().Get(reloadDeckQuery(
		d.API.BaseURL(),
		rules.BlackjackTable.AmountOfDecks(),
		rules.BlackjackTable.CardsPerDeck(),
	))

	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return errors.New("failed to fetch from external API")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if err = resp.Body.Close(); err != nil {
		return err
	}

	//parse body
	var jsonBody ReloadDeckBody

	if err = json.Unmarshal(body, &jsonBody); err != nil {
		return err
	}

	//write new deck into csv
	records := parsing.CardsToRecords(jsonBody.Cards)

	if err = d.CSV.Truncate(); err != nil {
		return err
	}

	if err = d.CSV.Writer().WriteAll(records); err != nil {
		return err
	}

	//refresh the file's offset
	if err = d.CSV.Return(); err != nil {
		return err
	}

	return nil
}

//GetHand retrieves the requested cards from the storage
func (d *deck) GetHand(cardIndexes []int) ([]model.Card, error) {
	r := d.CSV.Reader()

	//get the requested cards
	cards, err := readIndexedCards(cardIndexes, r)

	if err != nil {
		return cards, err
	}

	//refresh the file's offset
	if err = d.CSV.Return(); err != nil {
		return []model.Card{}, err
	}

	return cards, nil
}

func readIndexedCards(indexes []int, r *csv.Reader) ([]model.Card, error) {
	cards := []model.Card{}

	//dicard first record
	if _, err := r.Read(); err != nil {
		return cards, err
	}

	sort.Ints(indexes)

	var i int
	for _, ci := range indexes {
		for ; i <= ci; i++ {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if i == ci {
				cards = append(cards, parsing.RecordToCard(record))
			}
		}
	}

	return cards, nil
}
