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
	
			
}