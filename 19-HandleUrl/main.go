package main

import (
	"fmt"
	"net/url"
)

const urlString string = "https://www.google.com/learn?course=go&lesson=webrequest"

func main() {
	fmt.Println("Welocme to Url Handler")
	fmt.Println("URL is: ", urlString)
	result, err := url.Parse(urlString)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Query: ", result.RawQuery)

	qparams := result.Query()

	fmt.Printf("Query Params: %T \n", qparams)

	fmt.Println("Course: ", qparams["course"])

	for _, val := range qparams {
		fmt.Println("Param is : ", val)
	}

	partsofURL := &url.URL{
		Scheme:   "https",
		Host:     "www.google.com",
		Path:     "/learn",
		RawQuery: "course=reactJs&lesson=webrequest",
	}

	anotherURL := partsofURL.String()

	fmt.Println("Another URL: ", anotherURL)

}
