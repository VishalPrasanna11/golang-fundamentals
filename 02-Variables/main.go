package main

import (
	"fmt"
)

const Pi = 3.1415 //constant value of Pi // cannot be changed // public

func main() {
	var username string = "Vishal Prasanna"
	fmt.Println("Hello, from ", username)
	fmt.Printf("Username is of type: %T \n", username)

	var isLogged bool = false
	fmt.Println("Is the user Logged? ", isLogged)
	fmt.Printf("isLogged is of type: %T \n", isLogged)

	var smallval uint8 = 255
	fmt.Println("Un Signed Int Value", smallval)
	fmt.Printf("UnIsgned Int is of type: %T \n", smallval)

	var smallval1 uint16 = 256
	fmt.Println("Un Signed Int Value", smallval1)
	fmt.Printf("UnIsgned Int is of type: %T \n", smallval1)

	var smallFloot float64 = 256.545
	fmt.Println("Un Signed Int Value", smallFloot)
	fmt.Printf("UnIsgned Int is of type: %T \n", smallFloot)

	var somenumber int
	fmt.Println("Un Signed Int Value", somenumber)
	fmt.Printf("UnIsgned Int is of type: %T \n", somenumber)

	//implicit type
	var somevalue = "Hello"
	fmt.Println("Some Value", somevalue)
	fmt.Printf("Some value is of type: %T \n", somevalue)

	somevalue = "Hello, World!"
	fmt.Println("Some Value", somevalue)

	var sommeIntValue = 256
	fmt.Println("Some Int Value", sommeIntValue)
	fmt.Printf("Some Int Value is of type: %T \n", sommeIntValue)

	sommeIntValue = 3000

	//no var style

	numberofUser := 3000
	fmt.Println("Number of Users", numberofUser)
	fmt.Printf("Number of Users is of type: %T \n", numberofUser)

	Weight := 3000.00
	fmt.Println("Weight", Weight)
	fmt.Printf("Weight is of type: %T \n", Weight)

	fmt.Println("Value of Pi", Pi)
	fmt.Printf("Value of Pi is of type: %T \n", Pi)

}
