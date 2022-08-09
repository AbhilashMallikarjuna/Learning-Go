package main

// Functions in structs are called Methods
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
	abhilash.GetStatus()
	abhilash.SetEmail()

	fmt.Printf("User details are %+v\n", abhilash)
	// Here you can see that the orignal user email id is still unchanged
	// That is because while passing user to the function a copy of it is created.
	// Original object is not passed.
	// The modification happened on the copied object
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func (user User) GetStatus() {
	fmt.Println("Is user active: ", user.Status)
}

func (user User) SetEmail() {
	user.Email = "test@go.dev"
	fmt.Println("New Email is :", user.Email)
}
