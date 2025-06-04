// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/patriceteda/tiny-telematics-app/handlers"
// )

// func main() {
// 	r := mux.NewRouter()

// 	r.HandleFunc("/register", handlers.RegisterCar).Methods("POST")
// 	r.HandleFunc("/cars", handlers.ListCars).Methods("GET")

// 	fmt.Println("Server running at http://localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", r))
// }

package main

import (
"log"
"github.com/patriceteda/tiny-telematics-app/handlers"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
         log.Fatalf("Error loading config: %v", err)
	}
	handlers.Run(cfg)
}


