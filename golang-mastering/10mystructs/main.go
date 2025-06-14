package main

import (
	"fmt"
	"math/rand"
	// "time"
)


type User struct {
	Name    string
	Email   string
	Status  bool
	Age     int

}

func main() {
	patrice := User {
		Name:   "Patrice",
		Email:  "patrice@example.com",
		Status: true,
		Age:    33,
	}

	fmt.Printf("My name is %s, my email is %s, my status is %t and my age is %d\n", patrice.Name, patrice.Email, patrice.Status, patrice.Age)

	if patrice.Status {
		fmt.Println("User is active")
	} else {
		fmt.Println("User is inactive")
	}

	fmt.Println("==============================================")

	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount < 20 {
		result = "Moderate User"
	} else {
		result = "Admin User"
	}

	fmt.Println("User Type:", result)

	fmt.Println("==============================================")
	fmt.Println("Switch Case Example")

	// Switch case in golang
	// Lets create a simple dice game
	diceNumber := rand.Intn(6) + 1 // Random number between 1 and 6
	
	switch diceNumber {
case 1:
	fmt.Println("You rolled a 1, try again!")
case 2:
	fmt.Println("You rolled a 2, not bad!")
case 3:
	fmt.Println("You rolled a 3, good job!")
case 4:
	fmt.Println("You rolled a 4, great!")
case 5:
	fmt.Println("You rolled a 5, excellent!")
case 6:
	fmt.Println("You rolled a 6, amazing roll!")
}

		
}