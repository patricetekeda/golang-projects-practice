package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	// "os"
)


func main() {
	// PerformGetRequest()
	PerformPostRequest()

}


// PerformGetRequest makes a GET request to a specified URL and prints the response status and body.
// It uses the "net/http" package to send the request and read the response.
func PerformGetRequest () {
	const url = "https://jsonplaceholder.typicode.com/posts"

	resp, err := http.Get(url)
	if err != nil  {
		panic(err)
	}
	defer resp.Body.Close()
	
	fmt.Println("Response Status:", resp.Status)

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response Body:", string(content))
}


// PerformPostRequest makes a POST request to a specified URL with a JSON body.
// It uses the "net/http" package to send the request and read the response.
func PerformPostRequest() {
	const url = "https://jsonplaceholder.typicode.com/posts"

	// Create a new request body
	// This is a JSON body for the POST request
	resquestBody := strings.NewReader(
		`{
		"courseName": "Golang Mastering",
		"price": 0,
		"platform": "Udemy",
		"tags": ["golang", "programming", "backend"],
		"status": "active"
	    }`,
	)

	response, err := http.Post(url, "application/json", resquestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Response Status:", response.Status)

	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response Body:", string(content))
	fmt.Println("Post Request Completed Successfully")

}



