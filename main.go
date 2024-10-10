package main

import (
	"fmt"
	"log"
	freegamesscraper "main/freeGamesScraper"
	listscraper "main/listScraper"
)

func main() {
	listscraper.ParseGameList("listScraper/gameList.txt")
	freeGameName, isGameInList, err := freegamesscraper.CheckFreeGame()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(freeGameName, isGameInList, err)

}
