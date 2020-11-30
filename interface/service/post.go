package service

import (
	"github.com/JosueSdev/golang-bootcamp-2020/interface/model"

	"gorm.io/gorm"
)

type postService struct {
	db *gorm.DB
}

//PostService is an interface to communicate with the post's service
type PostService interface {
	GetAll() ([]*model.Post, error)
}

//NewPostService constructs a new PostService
func NewPostService(db *gorm.DB) PostService {
	return &postService{db}
}

//GetAll retrieves all posts from the database
func (ps *postService) GetAll() ([]*model.Post, error) {
	var posts []*model.Post

	result := ps.db.Find(&posts)

	if err := result.Error; err != nil {
		return nil, err
	}

	return posts, nil
}
