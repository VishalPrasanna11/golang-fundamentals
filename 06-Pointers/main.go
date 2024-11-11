package main

import "fmt"

func main() {

	fmt.Println("Welcome to Pointers in GoLang")

	myNumber := 10

	var ptr = &myNumber

	fmt.Println("My Number is: ", myNumber)
	fmt.Println("My Number's Address is: ", ptr)
	fmt.Println("My Number's Address is: ", *ptr)

	*ptr = *ptr * 10

	fmt.Println("My Number is: ", myNumber)
}
