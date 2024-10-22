package main

import (
	"fmt"
	freegamesscraper "main/freeGamesScraper"
	listscraper "main/listScraper"
)

func main() {
	go listscraper.ParseGameList("C:/IT/go/EpicGamesCollector/listScraper/gameList.txt")
	freeGameName, isGameInList, err := freegamesscraper.CheckFreeGame("C:/IT/go/EpicGamesCollector/listScraper/gameList.txt")
	if err != nil {
		/*fmt.Println("No good free games. Press Enter to exit...")
		fmt.Scanln()*/
		return
	}
	fmt.Println(freeGameName, isGameInList, err)
	fmt.Println("Program has finished. Press Enter to exit...")
	fmt.Scanln()
}
