package model

import "github.com/JosueSdev/golang-bootcamp-2020/domain/model"

//Post adapts the domain's post to a gorm aware environment
type Post struct {
	ID uint
	model.Post
}

//TableName defines the table to lookup in the database
func (Post) TableName() string {
	return "post"
}
