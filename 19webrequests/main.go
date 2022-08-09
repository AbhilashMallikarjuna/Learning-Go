package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web request to LCO")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of response is %T \n", response)

	//Caller's responsibility to close the connection
	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
