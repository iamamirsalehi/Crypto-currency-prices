package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type priceCrawler interface {
	sellPrice() prices
	buyPrice() prices
}

type prices struct {
	englishName string
	persianName string
	price       int
}

type website struct {
	englishName    string
	persianName    string
	baseWebsiteUrl string
	isSellPrice    bool
}

func (website *website) sellPrice() prices {
	website.isSellPrice = true
	return baseCrawler(website)
}

func (website *website) buyPrice() prices {
	website.isSellPrice = false
	return baseCrawler(website)
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

	prices.persianName = doc.Find("#RWPCS-usdt-table-buyers tr:nth-child(1) > td:nth-child(1)").First().Text()

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

	prices.persianName = doc.Find("#RWPCS-usdt-table-buyers tr:nth-child(1) > td:nth-child(1)").First().Text()

	doc.Find("#RWPCS-usdt-table-sellers tr:nth-child(1) > td:nth-child(2)").Each(func(i int, s *goquery.Selection) {
		BestSellPriceForClients := removeTomanFromText(s.Text())

		prices.price, err = strconv.Atoi(BestSellPriceForClients)
		checkErr(err)
	})

	return prices
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func baseCrawler(website *website) prices {
	var crawlType string

	baseUrl := "https://arzex.io/tether"

	netClient := GetClient()

	resp, err := netClient.Get(baseUrl)
	checkErr(err)
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	if website.isSellPrice == true {
		crawlType = "sellers"
	} else {
		crawlType = "buyers"
	}

	crawlerPriceByLink := doc.Find("#RWPCS-usdt-table-" + crawlType + " tr a[href^='" + website.baseWebsiteUrl + "']").Parent().Parent().Find("td:nth-child(2)").Text()

	if len(crawlerPriceByLink) == 0 {
		crawlerPriceByLink = doc.Find("#RWPCS-usdt-table-" + crawlType + " tr img[src^='/wp-content/plugins/arzexio/images/" + website.baseWebsiteUrl + "']").Parent().Parent().Find("td:nth-child(2)").Text()
	}

	price, err := strconv.Atoi(
		removeTomanFromText(
			crawlerPriceByLink,
		),
	)

	return prices{
		englishName: website.englishName,
		persianName: website.persianName,
		price:       price,
	}
}

func removeTomanFromText(text string) string {
	return strings.TrimSpace(strings.Replace(strings.Replace(text, "تومان", "", -1), ",", "", -1))
}
