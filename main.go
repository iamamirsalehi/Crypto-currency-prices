package main

import (
	"fmt"
	log "github.com/llimllib/loglevel"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/http"
	"os"
	"strings"
)

type Link struct {
	url string
	text string
	depth int
}

type HttpErrors struct {
	original string
}

func LinkReader(resp *http.Response) []Link{
	page := html.NewTokenizer(resp.Body)
	links := []Link{}

	var start *html.Token
	var text string

	for{
		_ := page.Next()
		token := page.Token()

		if token.Type == html.ErrorToken
	}
}
func main(){
	baseUrl := "https://arzex.io/tether/"
// 	"crypto/tls"
//	config := &tls.Config{
//		InsecureSkipVerify: true,
//	}
//
//	transport := &http.Transport{
//		TLSClientConfig: config,
//	}
//
//	netClient := &http.Client{
//		Transport: transport,
//	}
//	response, err := netClient.Get(baseUrl)
//	checkErr(err)

	//body, err := ioutil.ReadAll(response.Body)
	//defer response.Body.Close()
	//checkErr(err)

}

func checkErr(err error){
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}