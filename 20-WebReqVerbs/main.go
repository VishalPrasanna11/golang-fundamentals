package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to WebServer in GoLang")
	//GetRequest()
	//POSTRequest()
	POSTFORMRequest()

}

func GetRequest() {
	fmt.Println("GET Request")
	const url = "https://jsonplaceholder.typicode.com/posts"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Printf("Status Code: %d \n", response.StatusCode)
	fmt.Printf("Content Length: %d \n", response.ContentLength)
	fmt.Printf("Content Type: %s \n", response.Header.Get("Content-Type"))

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	//fmt.Println("Data: ", databytes)
	fmt.Println("Data: ", string(databytes))
}

func POSTRequest() {
	fmt.Println("POST Request")

	const url = "https://jsonplaceholder.typicode.com/posts"

	requestBody := strings.NewReader(`{
		"title": "foo",
		"body": "bar",
		"userId": 1
	}`)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Printf("Status Code: %d \n", response.StatusCode)
	fmt.Printf("Content Length: %d \n", response.ContentLength)
	fmt.Printf("Content Type: %s \n", response.Header.Get("Content-Type"))

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Data: ", string(databytes))
}

func POSTFORMRequest() {

	const myurl = "https://jsonplaceholder.typicode.com/posts"

	data := url.Values{}
	data.Add("title", "foo")
	data.Add("body", "bar")
	data.Add("userId", "1")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Printf("Status Code: %d \n", response.StatusCode)
	fmt.Printf("Content Length: %d \n", response.ContentLength)
	fmt.Printf("Content Type: %s \n", response.Header.Get("Content-Type"))

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Data: ", string(databytes))

}
