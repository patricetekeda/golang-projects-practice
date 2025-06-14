package main


import (
	"fmt"
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
}