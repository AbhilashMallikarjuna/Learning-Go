package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating,", input)

	//Adding 1 to your rating
	// Since the input will be in the form of string we need to convert it into number
	// For converting we use "strconv" package
	// But if we use only strconv package, we still get error because the "\n" will also be there in the input
	// So we need to trim the input string
	// For that we use "strings" package

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// Handling error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Adding 1 to your rating: ", numRating+1)
	}
}
