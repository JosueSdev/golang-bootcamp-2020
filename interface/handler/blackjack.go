package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/JosueSdev/golang-bootcamp-2020/interface/service"
	"github.com/JosueSdev/golang-bootcamp-2020/usecase/blackjack"
)

type blackjackHandler struct {
	Service service.Deck
}

//Blackjack provides http handlers for the blackjack game
type Blackjack interface {
	PutTable(w http.ResponseWriter, r *http.Request)
	GetGame(w http.ResponseWriter, r *http.Request)
}

//NewBlackjackHandler constructs a blackjack handler eith dependency injection
func NewBlackjackHandler(s service.Deck) Blackjack {
	return &blackjackHandler{s}
}

//PutTable fetches a new deck an uses it to write/replace the csv datastore
func (bj *blackjackHandler) PutTable(w http.ResponseWriter, r *http.Request) {
	if err := bj.Service.ReloadDeck(); err != nil {
		log.Printf("reloading deck: %s\n", err)
		HTTPProblem(w, http.StatusServiceUnavailable, "")
		return
	}

	response, err := json.Marshal(putTableResponse{
		Status: "ok",
	})

	if err != nil {
		log.Printf("parsing response: %s\n", err)
		HTTPProblem(w, http.StatusInternalServerError, "")
		return
	}

	if _, err = fmt.Fprint(w, string(response)); err != nil {
		log.Printf("responding: %s\n", err)
	}
}

//GetGame takes indices for cards and returns the result of the game
func (bj *blackjackHandler) GetGame(w http.ResponseWriter, r *http.Request) {
	//get and validate the request body
	requestBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("reading body: %s\n", err)
		HTTPProblem(w, http.StatusInternalServerError, "")
		return
	}

	var requestBody getGameBody

	if err = json.Unmarshal(requestBytes, &requestBody); err != nil {
		log.Printf("parsing request: %s\n", err)
		HTTPProblem(w, http.StatusBadRequest, err.Error())
		return
	}

	for _, ci := range requestBody.CardIndexes {
		if !blackjack.IsPositionInTableBounds(ci) {
			err = errors.New("out of bounds")
			log.Printf("parsing request: %s\n", err)
			HTTPProblem(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	//get cards from the service and play the game
	hand, err := bj.Service.GetHand(requestBody.CardIndexes)

	if err != nil {
		log.Printf("getting cards: %s\n", err)
		HTTPProblem(w, http.StatusInternalServerError, "")
		return
	}

	score, gameStatus, err := blackjack.CalculateScore(hand)

	if err != nil {
		log.Printf("calculating score: %s\n", err)
		HTTPProblem(w, http.StatusInternalServerError, "")
		return
	}

	//send back the game's result
	response, err := json.Marshal(getGameResponse{
		Score:      score,
		Hand:       hand,
		GameStatus: gameStatus,
	})

	if err != nil {
		log.Printf("parsing response: %s\n", err)
		HTTPProblem(w, http.StatusInternalServerError, "")
		return
	}

	if _, err = fmt.Fprint(w, string(response)); err != nil {
		log.Printf("responding: %s\n", err)
	}
}
