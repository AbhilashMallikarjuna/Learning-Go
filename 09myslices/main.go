package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to class on slices")

	// Initialization
	var fruitList = []string{"apples", "banana", "oranges"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	// append - It returns a new slice so we need to reassign to the parent slice
	fruitList = append(fruitList, "Grapes")
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:3]) // starts the slice from the i to j-1
	fmt.Println(fruitList)

	// Another way of initialization
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 444
	highScores[2] = 777
	highScores[3] = 555
	//highScores[4] = 666 this will give error bcz we have initialized for only 4 items

	highScores = append(highScores, 666, 777, 888)
	fmt.Println(highScores)

	sort.Ints(highScores) // inplace sorting. it will not return sorted slice
	fmt.Println(highScores)

	// Removing value from slices based on index
	var index int = 2
	// ... is used for slice unpacking
	highScores = append(highScores[:index], highScores[index+1:]...)
	fmt.Println(highScores)
}
