package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")
	// Initialization
	// it is compulsory to define the size of the array
	var fruitList [4]string

	//Declaration
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	//fruitList[2] = "Orange"
	fruitList[3] = "Grapes"
	fmt.Println(fruitList)
	fmt.Println("Length of fruitlist is", len(fruitList))

	// Even if the no. of elements in array are 3, the length will show the size of array i.e 4

	// Another way of initializing
	var students = [3]string{"Abhi", "Sindhu", "Sammi"}
	fmt.Println("Students are ", students)
}
