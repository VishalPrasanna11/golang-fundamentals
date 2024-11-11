package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to Maps in GoLang")

	languages := make(map[string]string)

	languages["en"] = "English"
	languages["es"] = "Spanish"
	languages["fr"] = "French"
	languages["de"] = "German"
	languages["it"] = "Italian"

	fmt.Println("Languages are: ", languages)
	fmt.Println(("EN stands for: "), languages["en"])

	delete(languages, "en")
	fmt.Println("Languages are: ", languages)

	for key, value := range languages {
		fmt.Println("Key: ", key, "Value: ", value)
	}
}
