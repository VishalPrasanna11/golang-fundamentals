package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Handle Time in GoLang")
	presentTime := time.Now()

	fmt.Println("Present Time is: ", presentTime)

	fmt.Println("Present Time is: ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(1998, time.February, 5, 18, 30, 0, 0, time.UTC)

	fmt.Println("Created Date is: ", createDate)

	fmt.Println("Created Date is: ", createDate.Format("01-02-2006 15:04:05 Monday"))
}
