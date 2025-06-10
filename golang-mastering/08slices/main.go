package main


import (
	"fmt"
)



func main() {

	// fruitsList := []string{"apple", "banana", "cherry", "date", "elderberry"}
	// fmt.Printf("Type of fruitslist: %T\n", fruitsList)

	// fruitsList = append(fruitsList, "fig")
	// fmt.Println("Fruits List after appending fig:", fruitsList)

	// newFruitsList := append(fruitsList[1:3])
	// fmt.Println("New Fruits List after slicing from index 1 to 3:", newFruitsList)

	// //making slices using make
	// slice1 := make([]int, 5)
	// slice1[0] = 1
	// slice1[1] = 2
	// slice1[2] = 3
	// slice1[3] = 4
	// slice1[4] = 5 // length 5, capacity 5

	// slice2 := make([]int, 5, 10) // length 5, capacity 10
	// slice2[0] = 6
	// slice2[1] = 7
	// slice2[2] = 8
	// slice2[3] = 9
	// slice2[4] = 10 // length 5, capacity 10

	// fmt.Println("Slice1 Length:", len(slice1), "Capacity:", cap(slice1))
	// fmt.Println("Slice2 Length:", len(slice2), "Capacity:", cap(slice2))


	// fmt.Println("Slice1:", slice1)
	// fmt.Println("Slice2:", slice2)

	studentsNames := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
	fmt.Println("Student Names:", studentsNames)
	
	vehicles := make([]string, 4, 9)
	vehicles[0] = "Car"
	vehicles[1] = "Bike"
	vehicles[2] = "Bus"
	vehicles[3] = "Truck"

	fmt.Println("Vehicles:", vehicles)
	fmt.Println("Length of Vehicles:", len(vehicles))

	var clothings []string = []string{"Shirt", "Pants", "Jacket", "Hat"}
	fmt.Println("Clothings:", clothings)
	fmt.Println("Length of Clothings:", len(clothings))

}