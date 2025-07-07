package models

// Author struct represents the author of a course
// It contains fields for the author's full name, email, and website.
type Author struct {
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Website string `json:"website"`
}



// Course struct represents a course in the system
// It contains fields for the course ID, name, price, and an embedded Author struct.
type Course struct {
	ID	string `json:"id"`
	CourseName string `json:"coursename"`
	Price int `json:"price"`
	Author *Author `json:"author"`
}