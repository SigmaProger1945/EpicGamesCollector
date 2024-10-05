package main

import (
	"fmt"
	"log"
	"main/db"
	freegamesscraper "main/freeGamesScraper"
	listscraper "main/listScraper"
)

func main() {
	Db, err := db.NewEditDb()
	if err != nil {
		panic(err)
	}
	listscraper.ParseGameList(Db)
	freeGameName, isGameInList, err := freegamesscraper.CheckFreeGame()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(freeGameName, isGameInList, err)

}
