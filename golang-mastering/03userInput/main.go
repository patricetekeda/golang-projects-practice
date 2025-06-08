package main

import (
	 "bufio"
     "os"
	 "fmt"
)


func main() {
	welcome := "Welcome to Go!"
	
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for your pizza: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating: ", input)

	fmt.Printf("The type of this rating is %T\n", input)


}