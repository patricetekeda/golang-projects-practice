package main

import (
	"fmt"
	"io"
	"os"
)


func main() {
	fmt.Println("Files in Go")

	content := "This is a sample content for the file."

	file, err := os.Create("./sample.txt")
	if err != nil {  
		panic(err)    
	}
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Println("File created successfully")
	fmt.Printf("File created with length: %d\n", length)
    defer file.Close()

	readFile("./sample.txt")

}

func readFile( fileName string) {
	content, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println("File Read Successfully")
	fmt.Println("File Content:", string(content))
	defer os.Remove(fileName) // Clean up the file after reading
	
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}


