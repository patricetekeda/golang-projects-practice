package main

import (
	"fmt"
)

type Course struct {
	CourseId   string
	CourseName string
}

// Method to check if Course is empty
func (c Course) IsEmpty() bool {
	if c.CourseId == "" && c.CourseName == "" {
		fmt.Println("Course is empty")
		return true	
	}
	fmt.Println("Course is not empty")
	return false

}

func main() {
	// Initialize a Course instance
	myCourse := Course{
		CourseId:   "CS101",
		CourseName: "Introduction to Computer Science",
	}

	// Check if it's empty and print result
	fmt.Println(myCourse.IsEmpty()) // Output: Course is not empty
	
}
