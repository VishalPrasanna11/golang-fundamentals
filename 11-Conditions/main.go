package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to Conditions in GoLang")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a Number: ")
	myNumber, _ := reader.ReadString('\n')

	myNumberInt, err := strconv.ParseInt(strings.TrimSpace(myNumber), 10, 64)

	if err != nil {
		fmt.Println("Error: ", err)
	} else if myNumberInt%2 == 0 {
		fmt.Println("My Number is Even")
	} else {
		fmt.Println("My Number is Odd")
	}

	newNumber := 10

	for newNumber > 0 {

		if newNumber%2 == 0 {
			fmt.Println("New Number is Even: ", newNumber)
		} else {
			fmt.Println("New Number is Odd: ", newNumber)
		}
		fmt.Println("New Number is: ", newNumber)
		newNumber--
	}
}
