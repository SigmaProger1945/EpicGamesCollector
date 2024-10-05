package listscraper

import (
	"log"
	"net/http"

	db "main/db"

	"github.com/PuerkitoBio/goquery"
	"gorm.io/gorm"
)

func ParseGameList(Db *db.EditDb) error {
	var listurl = "https://playthatgame.co.uk/?action=listbyplatform&platform=1"
	var gamelist db.GameList

	resp, err := http.Get(listurl)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	doc.Find("table tbody").Each(func(_ int, tr *goquery.Selection) {
		tr.Find("tr").Each(func(ix int, td *goquery.Selection) {
			t := td.Find("td").Eq(1)
			tres := string(t.Text())
			res := Db.Db.First(&gamelist, "game = ?", tres)
			if res.Error != gorm.ErrRecordNotFound {
				return
			}
			Db.AddGame(string(t.Text()))
		})
	})
	return nil
}
