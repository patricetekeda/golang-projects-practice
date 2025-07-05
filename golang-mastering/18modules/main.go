package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)




func main() {
	fmt.Println("Hello, Mod in Go!")
	greeter()
	r := mux.NewRouter()
	// Define routes
	r.HandleFunc("/", serveHome).Methods("GET")

	http.ListenAndServe(":8080", r)
	fmt.Println("Server is running on port 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
	// The above line is redundant since ListenAndServe is already called
	// but it can be used to log errors if the server fails to start.
	// If you want to log errors, you can use the following line instead:

}


func greeter() {
	fmt.Println(("Hey there!"))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	// Serve home page logic here
	w.Write([]byte("<h1>Welcome to the Home Page</h1>"))

}


