package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Course Model
type Course struct {
	ID         string `json:"id"`
	CourseName string `json:"coursename"`
	Price      int    `json:"price"`
}

// Fake database to store courses
var CourseDatabase = []Course{
	{
		ID:         "1",
		CourseName: "JavaScript Basics",
		Price:      100,
	},

	{
		ID:         "2",
		CourseName: "Go Programming Basics",
		Price:      150,
	},
	{
		ID:         "3",
		CourseName: "C# Programming Basics",
		Price:      120,
	},
}

func main() {
	fmt.Println("Hello, Mod in Go!")
	greeter()

	// Initialize the router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/api/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/api/courses/{id}", getCourseByID).Methods("GET")
	r.HandleFunc("/api/courses", createOneCourse).Methods("POST")

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func greeter() {
	fmt.Println("Hey there!")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	// Serve home page logic here
	w.Write([]byte("<h1>Welcome to the Home Page</h1>"))
	w.Write([]byte("<h1>This is basic API that handles very simple requests to a course databases</h1>"))
}

// Handler to get all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsondata, err := json.Marshal(CourseDatabase)
	if err != nil {
		http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsondata)
	fmt.Println("All courses retrieved successfully")

}

// Get a specific course by ID
func getCourseByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	courseID := params["id"]

	for _, course := range CourseDatabase {
		if course.ID == courseID {
			jsondata, err := json.Marshal(course)
			if err != nil {
				http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
				return
			}
			w.Write(jsondata)
			fmt.Println("Course retrieved successfully")
			return
		}
	}

	http.Error(w, "Course not found", http.StatusNotFound)
	fmt.Println("Course not found with ID:", courseID)

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check if the request body is empty (validation)
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send a request body")
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	// what about  - {}
    var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Validate course data
	if course.ID == "" || course.CourseName == "" || course.Price <= 0 {
		http.Error(w, "Invalid course data", http.StatusBadRequest)
		return
	}


	CourseDatabase = append(CourseDatabase, course)
	jsondata, err := json.Marshal(course)
	if err != nil {
		http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(jsondata)
	fmt.Println("New course created successfully")
}
