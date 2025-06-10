package main

import (
	"fmt"
)

func main() {

	age := 25
	name := "Patrice"

	namePointer := &name
	agePointer := &age
	
    println("Name:", name)
	println("Name Pointer:", namePointer)
	println("Name Value:", *namePointer)
	println("Age:", age)
	println("Age Pointer:", agePointer)
	println("Age Value:", *agePointer)

	println("================================")

	var pointer *int
	fmt.Println("value of the pointer:", pointer)

	println("================================")

	// Uing pointers in operations
	*agePointer = *agePointer + 1
	fmt.Println("New Age Value:", *agePointer)

}