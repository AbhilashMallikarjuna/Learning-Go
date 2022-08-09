package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=sdf456adf"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	// Parsing the URl
	result, _ := url.Parse(myurl)

	fmt.Println(result.Path)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	queryParams := result.Query()

	fmt.Printf("The type of query params are %T\n", queryParams)
	fmt.Println(queryParams["coursename"])

	partsOfURL := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "tutcss",
		RawPath: "user=abhilash"}

	newUrl := partsOfURL.String()
	fmt.Println(newUrl)
}
