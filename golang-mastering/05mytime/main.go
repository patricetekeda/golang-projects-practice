package main

import (
	"fmt"
	"time"
)


func main() {
	fmt.Println("Welcome to our time in Goland!")

	presentTime := time.Now()
	fmt.Println("Current time is:", presentTime.Format("01-02-2006 15:04:05 Monday"))
	
	// UTC time is universal time coordinated
	createdTime := time.Date(2023, time.October, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println("Created time is:", createdTime.Format("01-02-2006 15:04:05 Monday"))

}