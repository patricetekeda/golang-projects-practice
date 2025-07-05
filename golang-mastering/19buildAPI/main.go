package main

import (
	"encoding/json"
	"fmt"
)

// Goal is to build a simple API server in go
// Users will be using this API to access various courses
type Author struct {
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Website string `json:"website"`
}


type Course struct {
	ID	string `json:"id"`
	CourseName string `json:"coursename"`
	Price int `json:"price"`
	Author *Author `json:"author"`
}

// middleware
func (c *Course) IsEmpty() bool {
	// return true if the course is empty
	return c.ID == "" && c.CourseName == "" && c.Price == 0 
}


// Fake database will be used to store courses in a slice
var courseDatabase = []Course{}

func main() {
  // Initialize the course database with some data
  javascriptCourse := Course {
	ID: "1",
	CourseName: "JavaScript Basics",
	Price: 100,
	Author: &Author{
		FullName: "John Doe",
		Email: "john@example.com",
		Website: "https://johndoe.com",
	},
  }

  goCourse := Course {
	ID: "2",
	CourseName: "Go Programming Basics",
	Price: 150,
	Author: &Author{
		FullName: "Jane Smith",
		Email: "jane@example.com",
		Website: "https://janesmith.com",
	},
  }

  csharpCourse := Course {
	ID: "3",
	CourseName: "C# Programming Basics",
	Price: 120,
	Author: &Author{
		FullName: "Alice Johnson",
		Email: "alice@example.com",
		Website: "https://alicejohnson.com",
	},
  }

  // Add courses to the database
  courseDatabase = append(courseDatabase, javascriptCourse, goCourse, csharpCourse)

  fmt.Println("Course Database Initialized with Sample Data")

  // Pretty-print the courses as JSON
  prettyCourses, err := json.MarshalIndent(courseDatabase, "", "  ")
  if err != nil {
    fmt.Println("Error marshalling courses:", err)
    return
  }
  fmt.Println("Courses:")
  fmt.Println(string(prettyCourses))

}




