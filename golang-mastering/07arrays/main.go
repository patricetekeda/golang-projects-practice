package main


import (
	"fmt"
)


func main() {
	// Arrays in GO
	var numbers [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Numbers Array:", numbers)

	// remove an element from the array
	numbers[2] = 0 // setting the value to 0 to "remove" it
	fmt.Println("Numbers Array after removing element at index 2:", numbers)

	// length of the array
	lenght := len(numbers)
	fmt.Println("Length of the array:", lenght)

	// iterate over the array
	for index, value := range(numbers) {
		if value != 0 {
			fmt.Printf("Index: %d, Value: %d\n", index, value)
		}
			
	  }
	}



