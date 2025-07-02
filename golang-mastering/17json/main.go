package main

import (
	"fmt"
	"encoding/json"
)

type course struct {
	Name     string
	Price    int
	Platform string
	Password string
	Tags     []string
}

func main() {
	// Call the function to encode the course struct to JSON format
	// EncodeJSON()
	
	// You can also decode JSON data back to the course struct if needed
	DecodeJSON()
}

func EncodeJSON() {
	// This function will encode the course struct to JSON format
	programmingCourses := []course{
		{
			Name:     "Go Programming",
			Price:    100,
			Platform: "Udemy",
			Password: "golang123",
			Tags:     []string{"programming", "go", "backend"},
		},
		{
			Name:     "Python Programming",
			Price:    120,
			Platform: "Coursera",
			Password: "python123",
			Tags:     []string{"programming", "python", "data science"},
		},
	}

	jsonData, err := json.MarshalIndent(programmingCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println("JSON Data:", string(jsonData))
}

// consume the JSON data

func DecodeJSON() {
	// This function will decode JSON data back to the course struct
	
	jsonData := `[
	{
		"Name": "Go Programming",
		"Price": 100,
		"Platform": "Udemy",
		"Password": "golang123",
		"Tags": ["programming", "go", "backend"]
	},
	{
		"Name": "Python Programming",
		"Price": 120,
		"Platform": "Coursera",
		"Password": "python123",
		"Tags": ["programming", "python", "data science"]
	}
]`


	jsonDataFromWeb:= []byte(jsonData)

	var courses []course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &courses)
		fmt.Println("Decoded JSON Data:", courses)
	} else {
		fmt.Println("JSON is not valid")
	}

	var courseMap []map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &courseMap)
	fmt.Println("Decoded JSON Data as Map:", courseMap)

	for _, course := range courseMap {
		fmt.Printf("Name: %s, Price: %v, Platform: %v, Password: %v, Tags: %v\n",
			course["Name"], course["Price"], course["Platform"], course["Password"], course["Tags"])
	}
	fmt.Println("Decoded JSON Data as Map:", courseMap)
}
