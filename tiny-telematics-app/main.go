package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patriceteda/tiny-telematics-app/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.RegisterCar).Methods("POST")
	r.HandleFunc("/cars", handlers.ListCars).Methods("GET")

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}





