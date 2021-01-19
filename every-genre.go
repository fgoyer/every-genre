package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

var (
	genres []string
)

func main() {
	list, err := Genres()
	if err != nil {
		log.Fatal("There was a problem:", err)
	}

	fmt.Println("Number of Spotify Genres:", len(list))
}

func Genres() ([]string, error) {
	if len(genres) > 0 {
		return genres, nil
	}

	scraper := colly.NewCollector()
	scraper.OnHTML("div[preview_url]", func(element *colly.HTMLElement) {
		genre := strings.TrimSuffix(strings.TrimSpace(element.Text), "»")
		// Add genre to slice
		genres = append(genres, genre)
	})

	err := scraper.Visit("http://everynoise.com/engenremap.html")
	if err != nil {
		return nil, err
	}

	return genres, nil
}
