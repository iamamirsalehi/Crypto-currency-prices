package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type prices struct {
	name  string
	price int
}

func GetArzexBestBuyPrices() prices {
	baseUrl := "https://arzex.io/tether"

	prices := prices{}

	netClient := GetClient()

	resp, err := netClient.Get(baseUrl)
	checkErr(err)
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	prices.name = doc.Find("#RWPCS-usdt-table-buyers tr:nth-child(1) > td:nth-child(1)").First().Text()

	doc.Find("#RWPCS-usdt-table-buyers tr:nth-child(1) > td:nth-child(2)").Each(func(i int, s *goquery.Selection) {
		BestBuyPriceForClients := removeTomanFromText(s.Text())

		prices.price, err = strconv.Atoi(BestBuyPriceForClients)
		checkErr(err)
	})
	return prices
}

func GetArzexBestSellPrices() prices {
	baseUrl := "https://arzex.io/tether"

	prices := prices{}

	netClient := GetClient()

	resp, err := netClient.Get(baseUrl)
	checkErr(err)
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	prices.name = doc.Find("#RWPCS-usdt-table-buyers tr:nth-child(1) > td:nth-child(1)").First().Text()

	doc.Find("#RWPCS-usdt-table-sellers tr:nth-child(1) > td:nth-child(2)").Each(func(i int, s *goquery.Selection) {
		BestSellPriceForClients := removeTomanFromText(s.Text())

		prices.price, err = strconv.Atoi(BestSellPriceForClients)
		checkErr(err)
	})

	return prices
}

func GetNobitexSellPrices() prices{
	baseUrl := "https://arzex.io/tether"

	netClient := GetClient()

	resp, err := netClient.Get(baseUrl)
	checkErr(err)
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	price, err := strconv.Atoi(
		removeTomanFromText(
			doc.Find("#RWPCS-usdt-table-sellers tr a[href^='https://nobitex.ir/']").Parent().Parent().Find("td:nth-child(2)").Text(),
		),
	)

	nobitexSellPrice := prices{
		name:  "Nobitex",
		price: price,
	}

	return nobitexSellPrice
}

func GetNobitexBuyPrices() prices{
	baseUrl := "https://arzex.io/tether"

	netClient := GetClient()

	resp, err := netClient.Get(baseUrl)
	checkErr(err)
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	price, err := strconv.Atoi(
		removeTomanFromText(
			doc.Find("#RWPCS-usdt-table-buyers tr a[href^='https://nobitex.ir/']").Parent().Parent().Find("td:nth-child(2)").Text(),
		),
	)

	nobitexBuyPrice := prices{
		name:  "Nobitex",
		price: price,
	}

	return nobitexBuyPrice
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func removeTomanFromText(text string) string {
	return strings.TrimSpace(strings.Replace(strings.Replace(text, "تومان", "", -1), ",", "", -1))
}
