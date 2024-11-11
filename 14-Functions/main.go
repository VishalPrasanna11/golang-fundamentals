package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Functions in GoLang")
	greetings()
	greetingsTwo()

	result := add(10, 20)
	fmt.Println("Result is: ", result)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a Number A: ")
	numA, _ := reader.ReadString('\n')
	fmt.Println("Enter a Number B: ")
	numB, _ := reader.ReadString('\n')

	numA = strings.TrimSpace(numA)
	numB = strings.TrimSpace(numB)
	numAInt, _ := strconv.Atoi(numA)
	numBInt, _ := strconv.Atoi(numB)
	sum := add(numAInt, numBInt)
	fmt.Println("Sum is: ", sum)

}

func greetingsTwo() {
	fmt.Println("Hello GoLang")
}

func greetings() {
	fmt.Println("Namaste GoLang")
}

func add(a int, b int) int {
	return a + b
}
