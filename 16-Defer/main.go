package main

import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")

	defer fmt.Println("First Function")
	defer fmt.Println("Second Function")
	defer fmt.Println("Third Function")
	defer fmt.Println("Fourth Function")

	fmt.Println("Main Function")
	myDefer()

}

func myDefer() {

	for i := 0; i < 5; i++ {
		defer fmt.Println("Defer in myDefer Function", i)
	}
}
