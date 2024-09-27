package gamelistscraper

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Database struct {
	EditDb *db.EditDb
}

func ParseGameList() {
	var listurl = "https://playthatgame.co.uk/?action=listbyplatform&platform=1"
	resp, err := http.Get(listurl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("table tbody").Each(func(_ int, tr *goquery.Selection) {
		tr.Find("tr").Each(func(ix int, td *goquery.Selection) {
			t := td.Find("td").Eq(1)
			db.EditDb.AddGame(db.EditDb{}, string(t.Text()))
			fmt.Println(t.Text())

		})
	})
}
