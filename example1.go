package main

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

func main() {
	for i := 0; i < 5; i++ {
		webRequest(i)
	}
}

func webRequest(i int) {
	_, responseBody, _ := gorequest.New().Get("http://localhost:3000/").End()
	fmt.Printf("Request %v: %v\n", i, responseBody)
}
