package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome,"
	reader := bufio.NewReader(os.Stdin)

	// comma ok || comma err syntax

	input, _ := reader.ReadString('\n')
	fmt.Println(welcome, input)

	//	inputs are always stored as strings
}
