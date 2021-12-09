package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {

	urgArg := flag.String("url", "https://dir.bg", "the start url")
	flag.Parse()
	if urgArg == nil {
		log.Fatal("provide valid url")
	}

	startUrl , err := url.ParseRequestURI(*urgArg)
	if err != nil {
		fmt.Println("Please specify start page")
		os.Exit(1)
	}
	fmt.Println(startUrl)
}
