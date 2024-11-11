package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs in GoLang")

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var user1 User

	user1.FirstName = "Vishal"
	user1.LastName = "Prasanna"
	user1.Email = "user1@gmail.com"
	user1.Age = 25

	var user2 User

	user2.FirstName = "John"
	user2.LastName = "Doe"
	user2.Email = "user2@gmail.com"
	user2.Age = 30

	fmt.Println("User1 is: ", user1)
	fmt.Println("User2 is: ", user2)

}
