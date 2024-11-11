package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops in GoLang")

	// For Loop
	for i := 0; i < 5; i++ {
		fmt.Println("Value of i in for loop is: ", i)
	}

	// While Loop
	j := 0
	for j < 5 {
		fmt.Println("Value of j in while loop is: ", j)
		j++
	}

	// Infinite Loop
	k := 0
	for {
		fmt.Println("Value of k in infine loop  is: ", k)
		k++
		if k == 5 {
			break
		}
	}

	// Nested Loop
	for l := 0; l < 3; l++ {
		for m := 0; m < 3; m++ {
			fmt.Println("Value of l is: ", l, "Value of m is: ", m)
		}
	}

	// Looping through an Array

	myArray := [5]int{1, 2, 3, 4, 5}

	for n := 0; n < len(myArray); n++ {
		fmt.Println("Value of n is: ", myArray[n])
	}

	// Looping through an Array using Range
	daysOfWeek := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for index, day := range daysOfWeek {
		fmt.Println("Day: ", day, "Index: ", index)
	}

}
