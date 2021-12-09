package main

import (
	"crawler/crawl"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {

	urlArg := flag.String("url", "https://dir.bg", "the start url")
	flag.Parse()
	if urlArg == nil {
		log.Fatal("provide valid url")
	}

	startUrl , err := url.ParseRequestURI(*urlArg)
	if err != nil {
		fmt.Println("Please specify start page")
		os.Exit(1)
	}
	fmt.Println(startUrl)

	resp, err := http.Get(*urlArg)
	if err != nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()
	links := crawl.All(resp.Body, "img", "src")
	fmt.Println(links)
}
