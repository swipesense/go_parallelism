package main

import (
	"fmt"
	"sync"

	"github.com/parnurzeal/gorequest"
)

func main() {
	waitGroup := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go webRequest(i, waitGroup)
	}

	fmt.Println("Waiting")
	waitGroup.Wait()
}

func webRequest(i int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	_, responseBody, _ := gorequest.New().Get("http://localhost:3000/").End()
	fmt.Printf("Request %v: %v\n", i, responseBody)
}
