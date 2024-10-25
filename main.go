package main

import (
	"fmt"
	freegamesscraper "main/freeGamesScraper"
	listscraper "main/listScraper"
)

func main() {
	path := "gameList.txt"
	go listscraper.ParseGameList(path)
	freeGameName, isGameInList, err := freegamesscraper.CheckFreeGame(path)
	if err != nil {
		return
	}
	fmt.Println(freeGameName, isGameInList, err)
	fmt.Println("Program has finished. Press Enter to exit...")
	fmt.Scanln()
}
