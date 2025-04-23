package main

import (
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", GetUser).Methods("GET")


	http.ListenAndServe(":8080", r)

	
}

func GetUser(w http.ResponseWriter,r *http.Request){
	
	user := map[string]string{
		"id": "1",
		"name": "John Doe",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(user) 
	if err != nil {
		panic(err)
	}
	w.Write(jsonData)
}
