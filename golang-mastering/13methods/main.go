package main

import "fmt"


func main() {

	user := User{
		Name:   "John Doe",
		Email:  "john.doe@example.com",
		Status: true,
		Age:    30,
	}

	fmt.Println("User Details:")
	fmt.Println("Name:", user.Name)
	fmt.Println("Email:", user.Email)
	fmt.Println("Status:", user.Status)
	fmt.Println("Age:", user.Age)
	

	user.GetStatus()
	user.GetNewEmail()
	

}


type User struct {
	Name   string
	Email  string
	Status bool
	Age    int

}

func (u User) GetStatus() bool {
	fmt.Println("Is user active: ", u.Status)
	return u.Status
}

func (u User) GetNewEmail() {
	u.Email = "patrice@reeva.com"
	fmt.Println("New Email:", u.Email)
}



