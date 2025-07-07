package main

import (
	"encoding/json"
	"fmt"
	"github.com/patricetekeda/buildAPI/models/courses_model"

	
)

// middleware
func (c *courses_model.Course) IsEmpty() bool {
	// return true if the course is empty
	return c.ID == "" && c.CourseName == "" && c.Price == 0 
}


// Fake database will be used to store courses in a slice
var courseDatabase = []courses_model.Course{}

func main() {
  // Initialize the course database with some data
  javascriptCourse := courses_model.Course {
	ID: "1",
	CourseName: "JavaScript Basics",
	Price: 100,
	Author: &courses_model.Author{
		FullName: "John Doe",
		Email: "john@example.com",
		Website: "https://johndoe.com",
	},
  }

  goCourse := courses_model.Course {
	ID: "2",
	CourseName: "Go Programming Basics",
	Price: 150,
	Author: &courses_model.Author{
		FullName: "Jane Smith",
		Email: "jane@example.com",
		Website: "https://janesmith.com",
	},
  }

  csharpCourse := courses_model.Course {
	ID: "3",
	CourseName: "C# Programming Basics",
	Price: 120,
	Author: &courses_model.Author{
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




