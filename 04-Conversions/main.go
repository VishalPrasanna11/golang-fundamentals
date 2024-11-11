package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Your Pizza Shop: ")
	fmt.Println("Please give rating for our Pizza: ")
	rating, _ := reader.ReadString('\n')
	fmt.Println("Rating is ", rating)
	fmt.Printf("Rating is of type: %T \n", rating)
	fmt.Println("Adding +1 to the rating")
	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		numRating = numRating + 1
		fmt.Println("Rating is ", numRating)
	}

}
