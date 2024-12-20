package listscraper

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
)

func ParseGameList(listpath string) error {
	var listurl = "https://playthatgame.co.uk/?action=listbyplatform&platform=1"
	fiW, err := os.OpenFile(listpath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening FiW", err)
		return err
	}
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
	counter := 1
	doc.Find("table tbody").Each(func(_ int, tr *goquery.Selection) {
		tr.Find("tr").Each(func(ix int, td *goquery.Selection) {
			t := td.Find("td").Eq(1)
			game := FormatText(string(t.Text()))
			if IsGameInList(listpath, game) {
				return
			}
			line := fmt.Sprintf("%s\n", game)
			_, err := fiW.WriteString(line)
			if err != nil {
				fmt.Println("Failed to write to file:", err)
				return
			}
			counter++
		})
		//fmt.Println("Successful filewrite")
	})
	return nil
}

func IsGameInList(filepath string, title string) bool {
	fiR, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening FiR", err)
		return false
	}
	scanner := bufio.NewScanner(fiR)

	line := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), title) {
			return true
		}
		line++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("scan error", err)
		return false
	}
	return false
}
func FormatText(s string) string {
	return RemoveSpace(strings.ToLower(s))
}

func RemoveSpace(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}
	return string(rr)
}
