package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
	Email     string
	Age       int
}

func main() {
	fmt.Println("Welcome to Structs in GoLang")

	var user1 User

	user1.FirstName = "Vishal"
	user1.LastName = "Prasanna"
	user1.Email = "user1@gmail.com"
	user1.Age = 15

	var user2 User

	user2.FirstName = "John"
	user2.LastName = "Doe"
	user2.Email = "user2@gmail.com"
	user2.Age = 30

	//fmt.Println("User1 is: ", user1)
	fmt.Println("User 1 full name is: ", user1.fullName())
	fmt.Println("User 1 is Adult: ", user1.isAdult())
	fmt.Println("User 1's Email is: ", user1.getEmail())

	user1.Email = user1.newEmail()

	fmt.Println("User 1 email Update", user1.newEmail())

	fmt.Println("User 1's Email is: ", user1.getEmail())

	//fmt.Println("User2 is: ", user2)
	fmt.Println("User 2 full name is: ", user2.fullName())
	fmt.Println("User 2 is Adult: ", user2.isAdult())
	fmt.Println("User 2's Email is: ", user2.getEmail())

	// Method Declaration

}

func (u User) fullName() string {
	return u.FirstName + " " + u.LastName
}

// Method Declaration

func (u User) isAdult() bool {
	return u.Age >= 18
}

func (u User) getEmail() string {
	return u.Email
}

// Method Declaration

func (u User) newEmail() string {
	email := "new" + u.Email
	return email
}
