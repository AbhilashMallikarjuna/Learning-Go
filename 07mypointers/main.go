package main

import "fmt"

func main() {
	fmt.Println("Welcomt to class on pointers")

	// Initializing a pointer
	var ptr *int
	// Default value of a pointer is <nil>
	fmt.Println("Value of pointer is", ptr)

	myNumber := 23
	// & is used for referencing
	var myptr = &myNumber
	fmt.Println("Value of myPointer is", myptr)
	fmt.Println("Value of myPointer is", *myptr)

	*myptr = *myptr * 2
	fmt.Println("New value is", myNumber)

}
