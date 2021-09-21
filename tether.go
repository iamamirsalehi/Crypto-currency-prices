package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type prices struct {
	sellPrice int
	buyPrice  int
}

func GetArzexBestPrices() prices {
	baseUrl := "https://arzex.io/tether"

	arzexPrices := prices{}

	netClient := GetClient()

	resp, err := netClient.Get(baseUrl)
	checkErr(err)
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	doc.Find("#RWPCS-usdt-table-sellers tr:nth-child(1) > td:nth-child(2)").Each(func(i int, s *goquery.Selection) {
		BestSellPriceForClients := strings.TrimSpace(strings.Replace(strings.Replace(s.Text(), "تومان", "", -1), ",", "", -1))

		arzexPrices.sellPrice, err = strconv.Atoi(BestSellPriceForClients)
		checkErr(err)
	})

	doc.Find("#RWPCS-usdt-table-buyers tr:nth-child(1) > td:nth-child(2)").Each(func(i int, s *goquery.Selection) {
		BestBuyPriceForClients := strings.TrimSpace(strings.Replace(strings.Replace(s.Text(), "تومان", "", -1), ",", "", -1))

		arzexPrices.buyPrice, err = strconv.Atoi(BestBuyPriceForClients)
		checkErr(err)
	})
	return arzexPrices
}

func GetNobitexPrices() {
	baseUrl := "https://arzex.io/tether"

	//nobitexPrices := prices{}

	netClient := GetClient()

	resp, err := netClient.Get(baseUrl)
	checkErr(err)
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)
	fmt.Println("asdads")

	doc.Find("#RWPCS-usdt-table-sellers tr a[href^='https://nobitex.ir/']::parent + td:nth-child(2)").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
/*		BestSellPriceForClients := strings.TrimSpace(strings.Replace(strings.Replace(s.Text(), "تومان", "", -1), ",", "", -1))

		nobitexPrices.sellPrice, err = strconv.Atoi(BestSellPriceForClients)
		checkErr(err)*/
	})

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
