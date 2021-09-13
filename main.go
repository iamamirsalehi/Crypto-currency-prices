package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

func main(){
	c := colly.NewCollector()

	c.OnHTML("#RWPCS-usdt-table-sellers tr:nth-child(1) > td:nth-child(2)", func(e *colly.HTMLElement) {
		priceWithCurrency := e.Text

		BestSellPrice := strings.Replace(strings.Replace(priceWithCurrency,"تومان", "", 1), ",", "", 1)

		fmt.Println("Best seller: ", BestSellPrice)
	})

	c.OnHTML("#RWPCS-usdt-table-buyers tr:nth-child(1) > td:nth-child(2)", func(e *colly.HTMLElement) {
		priceWithCurrency := e.Text

		BestBuyPrice := strings.Replace(strings.Replace(priceWithCurrency,"تومان", "", 1), ",", "", 1)

		fmt.Println("Best buyer: ", BestBuyPrice)
	})

	c.Visit("https://arzex.io/tether/")
}