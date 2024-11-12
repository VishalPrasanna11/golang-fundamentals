package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"CourseName"`
	Price    int
	Platform string   `json:"Website"`
	Password string   `json:"-"`
	Tag      []string `json:"Tags, omitempty"`
}

func main() {
	fmt.Println("Welcome to Json in GoLang")
	//encodeJson()
	decodeJson()
}

func encodeJson() {

	myCourses := []course{
		{"Machine Learning", 299, "Coursera", "admin@123", []string{"ML", "AI"}},
		{"Deep Learning", 199, "Udemy", "root@123", []string{"DL", "AI"}},
		{"Artificial Intelligence", 399, "Edureka", "student@123", []string{"AI"}},
	}

	finalJson, err := json.MarshalIndent(myCourses, "", "\t")

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(string(finalJson))
	}

}

func decodeJson() {
	jsonString := []byte(`[
		{
			"CourseName": "Machine Learning",
			"Price": 299,
			"Website": "Coursera",
			"Tags": ["ML", "AI"]
		},
		{
			"CourseName": "Deep Learning",
			"Price": 199,
			"Website": "Udemy",
			"Tags": ["DL", "AI"]
		},
		{
			"CourseName": "Artificial Intelligence",
			"Price": 399,
			"Website": "Edureka",
			"Tags": ["AI"]
		}
	]`)

	var myCourses []course

	checkValid := json.Valid([]byte(jsonString))

	err := json.Unmarshal(jsonString, &myCourses)

	if err != nil {
		fmt.Println("Error: ", err)
	} else if checkValid {
		fmt.Printf("Decoded Json: %#v\n", myCourses)
	}

	//some key value pairs
	var result []map[string]interface{}
	err = json.Unmarshal(jsonString, &result)
	if err != nil {
		fmt.Println("Error decoding to map:", err)
		return
	}

	// Print each course as key-value pairs.
	for i, course := range result {
		fmt.Printf("Course %d:\n", i+1)
		for k, v := range course {
			fmt.Printf("  Key: %s, Value: %v\n", k, v)
		}
	}
}
