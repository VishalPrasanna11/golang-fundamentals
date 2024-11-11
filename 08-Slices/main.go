package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices in GoLang")

	// Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.
	var fruitList = []string{"Apple", "Tomato", "Banana"}
	fmt.Printf("Type of Fruit List is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Pineapple")
	fmt.Println("Fruit List is: ", fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println("Fruit List is: ", fruitList)

	fruitList = append(fruitList[:len(fruitList)-1])
	fmt.Println("Fruit List is: ", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 777 // will throw an error

	fmt.Println("High Scores: ", highScores)

	highScores = append(highScores, 555, 666, 777)
	fmt.Println("High Scores: ", highScores)

	//sorting a slice
	sort.Ints(highScores)

	fmt.Println("High Scores: ", highScores)

	// Sorting in decreasing order

	sort.Sort((sort.Reverse(sort.IntSlice(highScores))))
	fmt.Println("High Scores: ", highScores)

	//How to remove a value from a slice based on index

	var courses = []string{"react", "angular", "vue", "svelte"}

	fmt.Println("Courses: ", courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses: ", courses)

}
