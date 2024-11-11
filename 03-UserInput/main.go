package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welocme := "Hello, from Vishal Prasanna's GoLang Playground!"
	fmt.Println(welocme)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	// comma ok syntax
	username, _ := reader.ReadString('\n')

	fmt.Println("UserName is ", username)
}
