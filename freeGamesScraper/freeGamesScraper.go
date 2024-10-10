package freegamesscraper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	listscraper "main/listScraper"
	"net/http"
)

type FreeGamesPromotions struct {
	Data FreeGamesPromotionsData `json:"data"`
}
type FreeGamesPromotionsData struct {
	Catalog FreeGamesPromotionsCatalog `json:"Catalog"`
}
type FreeGamesPromotionsCatalog struct {
	SearchStore FreeGamesPromotionsSearchStore `json:"searchStore"`
}
type FreeGamesPromotionsSearchStore struct {
	Elements []FreeGamesPromotionsElements `json:"elements"`
}
type FreeGamesPromotionsElements struct {
	Title string                   `json:"title"`
	Price FreeGamesPromotionsPrice `json:"price"`
}
type FreeGamesPromotionsPrice struct {
	TotalPrice FreeGamesPromotionsTotalPrice `json:"totalPrice"`
}
type FreeGamesPromotionsTotalPrice struct {
	DiscountPrice int `json:"discountPrice"`
}

func CheckFreeGame() (string, bool, error) {
	var listurl = "https://store-site-backend-static-ipv4.ak.epicgames.com/freeGamesPromotions?locale=en-US&country=UA&allowCountries=UA"
	resp, err := http.Get(listurl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var freeGamesPromotions FreeGamesPromotions
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal([]byte(b), &freeGamesPromotions)
	if err != nil {
		log.Fatal(err)
	}
	elements := freeGamesPromotions.Data.Catalog.SearchStore.Elements
	if err != nil {
		return "", false, err
	}

	for _, element := range elements {
		if element.Price.TotalPrice.DiscountPrice == 0 {
			if listscraper.IsGameInList("listScraper/gameList.txt", element.Title) {
				return element.Title, true, nil
			}
		}
	}
	return "", false, fmt.Errorf("game not found")
}
