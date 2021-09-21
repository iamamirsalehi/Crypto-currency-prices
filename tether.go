package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type prices struct {
	englishName string
	persianName string
	price       int
}

type crawlType struct {
	englishName    string
	persianName    string
	baseWebsiteUrl string
	isSellPrice    bool
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

func GetNobitexSellPrices() prices {
	buyPriceCrawler := crawlType{
		englishName:    "Nobitex",
		persianName:    "نوبیتکس",
		baseWebsiteUrl: "https://nobitex.ir/",
		isSellPrice:    true,
	}

	return base(buyPriceCrawler)
}

func GetNobitexBuyPrices() prices {
	sellPriceCrawler := crawlType{
		englishName:    "Nobitex",
		persianName:    "نوبیتکس",
		baseWebsiteUrl: "https://nobitex.ir/",
		isSellPrice:    false,
	}
	return base(sellPriceCrawler)
}

func GetPhinixSellPrices() prices {
	sellPriceCrawler := crawlType{
		englishName:    "Phinix",
		persianName:    "فینیکس",
		baseWebsiteUrl: "https://phinix.ir/",
		isSellPrice:    true,
	}
	return base(sellPriceCrawler)
}

func GetPhinixBuyPrices() prices {
	sellPriceCrawler := crawlType{
		englishName:    "Phinix",
		persianName:    "فینیکس",
		baseWebsiteUrl: "https://phinix.ir/",
		isSellPrice:    false,
	}
	return base(sellPriceCrawler)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func base(crawlerInfo crawlType) prices {
	var crawlType string

	baseUrl := "https://arzex.io/tether"

	netClient := GetClient()

	resp, err := netClient.Get(baseUrl)
	checkErr(err)
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	if crawlerInfo.isSellPrice == true {
		crawlType = "sellers"
	} else {
		crawlType = "buyers"
	}

	price, err := strconv.Atoi(
		removeTomanFromText(
			doc.Find("#RWPCS-usdt-table-" + crawlType + " tr a[href^='" + crawlerInfo.baseWebsiteUrl + "']").Parent().Parent().Find("td:nth-child(2)").Text(),
		),
	)

	return prices{
		englishName: crawlerInfo.englishName,
		persianName: crawlerInfo.persianName,
		price:       price,
	}
}

func removeTomanFromText(text string) string {
	return strings.TrimSpace(strings.Replace(strings.Replace(text, "تومان", "", -1), ",", "", -1))
}
