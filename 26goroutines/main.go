package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex
var signals = []string{"test"}

func main() {
	//go greeter("Hello")
	//greeter("World")
	webistelist := []string{

		"https://google.com",
		"https://twitter.com",
		"https://fb.com",
		"https://lco.dev",
		"https://go.dev",
		"https://abhilashmallikarjuna.me",
	}

	for _, site := range webistelist {
		go getStatusCode(site)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

//	func greeter(s string) {
//		for i := 0; i < 5; i++ {
//			time.Sleep(5 * time.Millisecond)
//			fmt.Println(s)
//		}
//	}
func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status for %s website\n", res.StatusCode, endpoint)
	}
}
