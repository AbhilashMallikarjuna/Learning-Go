package main

import "fmt"

func main() {
	fmt.Println("Functions in golang")
	greeter()

	result := adder(4, 5)
	fmt.Println("Result of two nums is", result)

	sum, msg := proAdder(1, 2, 3, 4)
	fmt.Println(msg, sum)
}

func adder(numOne int, numTwo int) int {
	return numTwo + numOne
}

func proAdder(nums ...int) (int, string) {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total, "Total sum is"

}

func greeter() {
	fmt.Println("Namasthe from Golang")
}

// It is not allowed to write functions inside functions
