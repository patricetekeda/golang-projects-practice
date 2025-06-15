package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Defer one")
	defer fmt.Println("Defer two")
	defer fmt.Println("Defer three")
    fmt.Println("Defer in Go")	

	myDefer()
}


func myDefer () {
	for index:= 0; index < 5; index++ {
		defer fmt.Println(index)
	}
}