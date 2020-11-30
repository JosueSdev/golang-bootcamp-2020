package model

import "time"

//Post is a storage-agnostic definition of a post for the board
type Post struct {
	Message string
	History int
	Time    time.Time
}

//Get implements PostGetter for Post
func (p Post) Get() Post {
	return p
}
