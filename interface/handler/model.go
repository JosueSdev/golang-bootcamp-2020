package handler

import "github.com/JosueSdev/golang-bootcamp-2020/domain/model"

type getGameBody struct {
	CardIndexes []int `json:"cards"`
}

type getGameResponse struct {
	Score      int          `json:"score"`
	Hand       []model.Card `json:"hand"`
	GameStatus string       `json:"game_status"`
}

type putTableResponse struct {
	Status string `json:"status"`
}
