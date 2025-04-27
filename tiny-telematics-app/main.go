package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/patriceteda/tiny-telematics-app/handlers"
)

func main() {
    http.HandleFunc("/register", handlers.RegisterCar)
    http.HandleFunc("/cars", handlers.ListCars)

    fmt.Println("Server running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8081", nil))
}
