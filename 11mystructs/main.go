package main

import "fmt"

func main() {
	fmt.Println("Structs in GoLang")
	// NO INHERITANCE
	// NO SUPER OR PARENT CONCEPT

	abhilash := User{"Abhilash", 25, "abhi@go.dev", true}
	fmt.Println(abhilash)
	// For detailed output
	fmt.Printf("User details are %+v\n", abhilash)
	// For specific details
	fmt.Printf("User name is %v and the age is %v\n", abhilash.Name, abhilash.Age)
}

type User struct {
	Name    string
	Age     int
	email   string
	verfied bool
}
