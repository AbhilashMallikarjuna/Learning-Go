package main

import "fmt"

const LoginToken = "sdfahsdfadfj"

// Here the LoginToken is public since the first letter is capital

func main() {
	var username string = "Abhilash"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)
	// Ranges of Numeric types
	//https://go.dev/ref/spec#Numeric_types

	var smallFloat float32 = 255.45544678634527
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//	The precision after decimal point is more for float64 than float32

	// Default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// Implicit types
	// If we don't specify the type, then the lexer automatically puts by seeing the variable contents
	var website = "google.com"
	fmt.Println(website)

	// No var style using Walrus operator
	// Walrus operator can only be used inside methods. Not to be used in global scopes
	numberOfUsers := 500
	fmt.Printf("Variable is of type: %T \n", numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
