package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video")
	PerformGetRequest()
	PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest(){
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)


	if err != nil{
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code is:", response.StatusCode)
	fmt.Println("Content length is:",response.ContentLength)

	
	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is:",byteCount)
	fmt.Println(responseString.String())


	// fmt.Println(string(content))


}

func PerformPostJsonRequest(){
	const myurl = "http://localhost:8000/post"

	payload := strings.NewReader(`
	{
		"name" : "Abhilash",
		"age" : "25",
		"email" : "abhi@go.dev"
	}
	`)
	response, err := http.Post(myurl, "application/json", payload)

	if err != nil{
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}

func PerformPostFormRequest(){
	const myurl = "http://localhost:8000/post"

	formdata := url.Values{}
	formdata.Add("first name","Abhilash")
	formdata.Add("last name","Mallikarjuna")
	formdata.Add("email","abhi@microsoft.com")

	response, err := http.PostForm(myurl, formdata)
	
	if err != nil{
		panic(err)
	}
	defer response.Body.Close()

	content , _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}