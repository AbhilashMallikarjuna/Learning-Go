package main

import (
	"fmt"
	"time"
)

// https://pkg.go.dev/time
func main() {
	fmt.Println("----------------------------------------")
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	// The string below should be written in the exact same manner. No modifications to be done
	// Other wise it will not give desired output
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.December, 10, 5, 10, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
