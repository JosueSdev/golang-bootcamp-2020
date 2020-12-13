package datastore

import (
	"encoding/csv"
	"os"

	"github.com/JosueSdev/golang-bootcamp-2020/config"
)

type csvDeck struct {
	file *os.File
}

//CSVDeck allows to interface with the csv file of the deck
type CSVDeck interface {
	Reader() *csv.Reader
	Writer() *csv.Writer
	Return() error
	Truncate() error
	Close() error
}

//NewCSVDeck opens a file and returns a CSVDeck
func NewCSVDeck() (CSVDeck, error) {
	f, err := os.OpenFile(config.DeckFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return nil, err
	}

	return &csvDeck{f}, nil
}

//Reader returns a csv reader
func (deck *csvDeck) Reader() *csv.Reader {
	return csv.NewReader(deck.file)
}

//Writer returns a csv writer
func (deck *csvDeck) Writer() *csv.Writer {
	return csv.NewWriter(deck.file)
}

//Close closes the internal file
func (deck *csvDeck) Close() error {
	return deck.file.Close()
}

//Return returns the cursor to the begining of the file
func (deck *csvDeck) Return() error {
	_, err := deck.file.Seek(0, 0)

	if err != nil {
		return err
	}

	return nil
}

//Truncate resets the csv file
func (deck *csvDeck) Truncate() error {
	err := deck.file.Truncate(0)

	if err != nil {
		return err
	}

	_, err = deck.file.Seek(0, 0)

	if err != nil {
		return err
	}

	return nil
}
