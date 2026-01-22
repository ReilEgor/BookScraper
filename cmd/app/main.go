package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type Book struct {
	Title string `json:"title"`
	Price string `json:"price"`
}

func main() {
	c := colly.NewCollector()

	var books []Book

	c.OnHTML(".product_pod", func(e *colly.HTMLElement) {
		bookInfo := Book{
			Title: e.ChildAttr("h3 a", "title"),
			Price: e.ChildText(".price_color"),
		}
		books = append(books, bookInfo)
	})
	c.OnScraped(func(r *colly.Response) {
		data, err := json.MarshalIndent(books, "", "    ")

		if err != nil {
			fmt.Println("Error marshalling data:", err)
			return
		}
		err = os.WriteFile("books.json", data, 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
	})
	//TODO: Add it in env variables
	c.Visit("https://books.toscrape.com/")

}
