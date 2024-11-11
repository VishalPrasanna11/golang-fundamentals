package main

import (
	"fmt"
	"time"

	"math/rand"
)

func main() {
	fmt.Println("Welcome to Switch in GoLang")

	rand.NewSource(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of Dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice Value is 1")
	case 2:
		fmt.Println("Dice Value is 2")
		// fallthrough keyword is used within a switch statement to allow the execution to "fall through" to the next case,
		//even if that case is not explicitly matched
		fallthrough
	case 3:
		fmt.Println("Dice Value is 3")
		fallthrough
	case 4:
		fmt.Println("Dice Value is 4")
	case 5:
		fmt.Println("Dice Value is 5")
	case 6:
		fmt.Println("Dice Value is 6")
	default:
		fmt.Println("Invalid Dice Value")

	}

}
