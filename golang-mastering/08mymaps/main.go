package main


import (
	"fmt"
)



func main() {
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"
	languages["TS"] = "TypeScript"

	fmt.Println("Languages Map:", languages)

	for key, value := range(languages) {
		fmt.Println("Key:", key, "Value:", value)
	}



}