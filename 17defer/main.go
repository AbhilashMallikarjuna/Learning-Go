package main

import "fmt"

// Defer functions will be executed before the return statement of the surronding functions
// In case of multiple defer statements , the functions will be executed in LIFO manner
func main() {
	defer fmt.Println("Statement 2")
	defer fmt.Println("Statement 3")
	fmt.Println("Statement 1")
	myDefer()
}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
