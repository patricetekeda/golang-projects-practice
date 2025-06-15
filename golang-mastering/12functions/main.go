package main


import (
	"fmt"
)



func main() {
	fmt.Println("functions in Go")

	greeter()   // Calling the greeter function

	adderResult := adder(5, 10) // Calling the adder function with two integers
	fmt.Println("The result of the adder function is:", adderResult)

	proAdderResult, proAdderMessage := proAdder(1, 2, 3, 4, 5) // Calling the proAdder function with a variable number of integers
	fmt.Println("The result of the proAdder function is:", proAdderResult)
	fmt.Println("Message from proAdder function:", proAdderMessage)

}


func greeter() {
	fmt.Println("Hello from the greeter function")
}

func adder(a int, b int) int {
	// This function takes two integers and returns their sum
	return a + b
}

func proAdder(values ...int) (int, string) {
	// This function takes a variable number of integers and returns their sum
	// values is a slice of integers
	// We can use the range keyword to iterate over the slice

	total := 0
	for _, value := range(values) {
		total += value
	}

	return total, "This is a variable number of integers function" // Returning the total and a message

}

