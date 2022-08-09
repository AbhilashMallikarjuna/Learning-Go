package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")
	days := []string{"Sun", "mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	// First way of writing for loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	fmt.Println("---------------")
	//Second way
	for i := range days {
		fmt.Println(days[i])
	}
	fmt.Println("---------------")

	// For each kind of loop
	for index, day := range days {
		fmt.Printf("Index is %v and day is %v\n", index, day)
	}
	fmt.Println("---------------")

	// While loop using for keyword
	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 2 {
			goto tag
		}
		fmt.Println("Rougue value is", rougueValue)
		rougueValue += 1
	}
tag:
	fmt.Println("Jumping to the tag")
}
