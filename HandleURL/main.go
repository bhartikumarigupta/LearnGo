package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("We are learning handle url in Go")
	myurl := "https://jsonplaceholder.typicode.com/"
	parseURL, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Printf("type of parseURL is %T\n", parseURL)
	fmt.Println("Scheme: ", parseURL.Scheme)
	fmt.Println("Host: ", parseURL.Host)
	fmt.Println("Path: ", parseURL.Path)
	fmt.Println("RawQuery: ", parseURL.RawQuery)
	fmt.Println("Fragment: ", parseURL.Fragment)
	// modify the URL
	parseURL.Path = "/path/post1/1"
	fmt.Println("Modified URL: ", parseURL)
	fmt.Println("Modified Path: ", parseURL.Path)
	newurl := parseURL.String()
	fmt.Println("New URL: ", newurl)

}
