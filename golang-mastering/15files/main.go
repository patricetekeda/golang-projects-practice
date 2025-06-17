package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)


func main() {
	fmt.Println("Files in Go")

	content := "This is a sample content for the file."

	file, err := os.Create("./sample.txt")
	CheckError(err)
	length, err := io.WriteString(file, content)
	CheckError(err)
	fmt.Println("File created successfully")
	fmt.Printf("File created with length: %d\n", length)
    defer file.Close()

	// readFile("./sample.txt")
	// GetWebcontent()
	ParseUrl()

}

func readFile( fileName string) {
	content, err := os.ReadFile(fileName)
	CheckError(err)
	fmt.Println("File Read Successfully")
	fmt.Println("File Content:", string(content))
	defer os.Remove(fileName) // Clean up the file after reading
	
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// Handling web request in golang
const postsURL string = "https://jsonplaceholder.typicode.com/posts"

func GetWebcontent() {
	resp, err := http.Get(postsURL)
	CheckError(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	CheckError(err)
	fmt.Println("Web Content Retrieved Successfully")
	fmt.Println("Web Content:", string(body))

}

// Handling url in golang
// url with query string or parameters
const url1 string = "https://jsonplaceholder.typicode.com/posts?userId=1"

func ParseUrl() {
	res, err := url.Parse(url1)
	CheckError(err)
	fmt.Println("Parsed URL:", res)
	fmt.Println("Scheme:", res.Scheme)
	fmt.Println("Host:", res.Host)
	fmt.Println("Path:", res.Path)
	fmt.Println("Query:", res.Query())
	fmt.Println("User ID:", res.Query().Get("userId"))
	fmt.Println("Fragment:", res.Fragment)
	fmt.Println("Raw Query:", res.RawQuery)
	fmt.Println("Raw Path:", res.RawPath)
	fmt.Println("String Representation:", res.String())
	fmt.Println("Host with Port:", res.Hostname(), ":", res.Port())
	fmt.Println("URL String:", res.String())
	fmt.Println("URL String with Query:", res.String()+"?"+res.RawQuery)
	fmt.Println("URL String with Fragment:", res.String()+"#"+res.Fragment)
	fmt.Println("URL String with Path:", res.String()+res.Path)

	fmt.Println("gettig query parameters")
	qparams := res.Query()
	for key, value := range qparams {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
	
}









