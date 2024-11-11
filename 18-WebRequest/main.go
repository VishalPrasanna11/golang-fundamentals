package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const url = "https://jsonplaceholder.typicode.com/posts"

func main() {
	fmt.Println("Welcome to Web Request in GoLang")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response: of Type %T \n ", response)

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	file, err := os.Create("./httplog.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("Length: ", length)
	//fmt.Println("Data: ", string(databytes))
	response.Body.Close()

}
