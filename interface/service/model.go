package service

import "github.com/JosueSdev/golang-bootcamp-2020/domain/model"

//ReloadDeckBody maps the json body of a /draw request
type ReloadDeckBody struct {
	Cards []model.Card `json:"cards"`
}
