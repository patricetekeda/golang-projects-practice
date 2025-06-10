package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
	fmt.Println("Welcome my pizza shop!")
	fmt.Println("Enter a rating between 1 and 5:")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating:", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {		
		fmt.Println("Error converting input to number:", err)
		return
	} else {
		fmt.Println("Your rating as a number is:", numRating + 1)
	}

}


