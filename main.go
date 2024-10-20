package main

import (
	"fmt"
	freegamesscraper "main/freeGamesScraper"
	listscraper "main/listScraper"
)

func main() {
	listscraper.ParseGameList("listScraper/gameList.txt")
	freeGameName, isGameInList, err := freegamesscraper.CheckFreeGame()
	if err != nil {
		fmt.Println("No good free games. Press Enter to exit...")
		fmt.Scanln()

	}
	fmt.Println(freeGameName, isGameInList, err)
	fmt.Println("Program has finished. Press Enter to exit...")
	fmt.Scanln()

}
