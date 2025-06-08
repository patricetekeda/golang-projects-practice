package main

import "fmt"





func main() {
	var username string = "Patrice"
	var age int = 33
	var isLoggedIn bool = false
	var smallValue uint8 = 255
	var smallFloat float32 = 255.5

	fmt.Println(username)
	fmt.Printf("Hello %s, welcome to Go!\n", username)
	fmt.Printf("I am %d years old.\n", age)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", smallValue)
	fmt.Printf("Variable is of type: %T\n", smallFloat)
	
}