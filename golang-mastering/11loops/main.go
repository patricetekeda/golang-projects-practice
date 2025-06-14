package main


import (
	"fmt"
)

// This is a simple for loop example in Go
func main() {

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for index, value := range(days) {

			fmt.Printf("This is day %d, which is a: %s\n", index+1, value)

	}

	fmt.Println("Another for loop syntax")

	for index:= 0; index < len(days); index++ {
		
		fmt.Printf("This is day %d, which is a: %s\n", index+1, days[index])
	}


	rogueValue := 1

	for rogueValue < 10 {
		fmt.Println("Rogue value is:", rogueValue)
		rogueValue++
		
		if rogueValue == 5 {
			fmt.Println("Rogue value is 5, breaking the loop")
			break
		} 
		
	}

}