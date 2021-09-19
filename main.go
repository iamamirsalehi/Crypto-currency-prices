package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	baseUrl := "https://arzex.io/tether"

	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: config,
	}

	netClient := &http.Client{
		Transport: transport,
	}

	resp, err := netClient.Get(baseUrl)
	checkErr(err)
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	doc.Find("#RWPCS-usdt-table-sellers tr:nth-child(1) > td:nth-child(2)").Each(func(i int, s *goquery.Selection) {
		BestSellPriceForClients := strings.Replace(strings.Replace(s.Text(), "تومان", "", -1), ",", "", -1)

		fmt.Println(BestSellPriceForClients)
	})

	doc.Find("#RWPCS-usdt-table-buyers tr:nth-child(1) > td:nth-child(2)").Each(func(i int, s *goquery.Selection) {
		BestBuyPriceForClients := strings.Replace(strings.Replace(s.Text(), "تومان", "", -1), ",", "", -1)

		fmt.Println(BestBuyPriceForClients)
	})

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
