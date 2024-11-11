package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Code;

	fmt.Println("Hello, from Vishal Prasanna's GoLang Playground!")
	content := "This is an example of writing to a file in GoLang from Vishal."

	file, err := os.Create("./myFileLog.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	length, err := io.WriteString(file, content)

	if err != nil {
		fmt.Println("Error: ", err)
		file.Close()
		return
	}
	fmt.Println("Length: ", length)
	defer file.Close()
	readFile("./myFileLog.txt")
}

func readFile(fileName string) {
	databyte, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)

	}
	fmt.Println("The text data is \n", string(databyte))

}
