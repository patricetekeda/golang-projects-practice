package main

import (
	"encoding/json"
	"fmt"
	"github.com/patricetekeda/golang-projects-practice/golang-mastering/buildAPI/models"
)

// Middleware
func IsCourseEmpty(c *models.Course) bool {
	// Return true if the course is empty
	return c.ID == "" && c.CourseName == "" && c.Price == 0
}

// Fake database to store courses in a slice
var courseDatabase = []models.Course{}

func main() {
	// Initialize the course database with some data
	javascriptCourse := models.Course{
		ID:         "1",
		CourseName: "JavaScript Basics",
		Price:      100,
		Author: &models.Author{
			FullName: "John Doe",
			Email:    "john@example.com",
			Website:  "https://johndoe.com",
		},
	}

	goCourse := models.Course{
		ID:         "2",
		CourseName: "Go Programming Basics",
		Price:      150,
		Author: &models.Author{
			FullName: "Jane Smith",
			Email:    "jane@example.com",
			Website:  "https://janesmith.com",
		},
	}

	csharpCourse := models.Course{
		ID:         "3",
		CourseName: "C# Programming Basics",
		Price:      120,
		Author: &models.Author{
			FullName: "Alice Johnson",
			Email:    "alice@example.com",
			Website:  "https://alicejohnson.com",
		},
	}

	// Add the courses to the database
	courseDatabase = append(courseDatabase, javascriptCourse, goCourse, csharpCourse)

	// Print the course database as JSON
	result, err := json.MarshalIndent(courseDatabase, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(result))
}
