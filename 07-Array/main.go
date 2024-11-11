package main

import "fmt"

func main() {

	fmt.Println("Welcome to Arrays in GoLang")

	// Array Declaration
	var myArray [5]int

	// Array Initialization
	myArray[1] = 10
	myArray[3] = 20

	fmt.Println("My Array is: ", myArray)
	fmt.Println("My Array's Length is: ", len(myArray))

	// Array Declaration and Initialization
	myArray2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("My Array2 is: ", myArray2)
	fmt.Println("My Array2's Length is: ", len(myArray2))

	// Array Declaration and Initialization
	myArray3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("My Array3 is: ", myArray3)
	fmt.Println("My Array3's Length is: ", len(myArray3))

	// Array Declaration and Initialization
	myArray4 := [5]string{"Vishal", "Prasanna", "GoLang", "Playground"}
	fmt.Println("My Array4 is: ", myArray4)
	fmt.Println("My Array4's Length is: ", len(myArray4))

}
