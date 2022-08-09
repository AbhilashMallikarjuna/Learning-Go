package main

import "fmt"

func main() {
	fmt.Println("If Else in golang")
	loginCount := 10

	if loginCount < 10 {
		fmt.Println("Normal user")
	} else if loginCount == 10 {
		fmt.Println("Potential convert")
	} else {
		fmt.Println("Power User")
	}

}
