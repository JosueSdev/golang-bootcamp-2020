package main

import (
	"fmt"

	"github.com/JosueSdev/golang-bootcamp-2020/infrastructure/datastore"
	"github.com/JosueSdev/golang-bootcamp-2020/interface/handler"
)

func main() {
	db, err := datastore.NewMysqlDB()

	if err != nil {
		panic(err)
	}

	postHandler := handler.NewPostHandler(db)

	if err = postHandler.GetAll(); err != nil {
		fmt.Println(err)
	}
}
