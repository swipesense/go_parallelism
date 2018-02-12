package main

import (
	"fmt"
	"sync"

	"github.com/parnurzeal/gorequest"
)

func main() {
	numWorkers := 5
	waitGroup := new(sync.WaitGroup)
	jobs := make(chan int, numWorkers)

	for workerID := 0; workerID < numWorkers; workerID++ {
		waitGroup.Add(1)
		go webWorker(workerID, jobs, waitGroup)
	}

	numJobs := 10
	for jobID := 0; jobID < numJobs; jobID++ {
		fmt.Printf("Submitting job %v\n", jobID)
		jobs <- jobID
	}
	close(jobs)

	fmt.Println("Waiting")
	waitGroup.Wait()
}

func webWorker(workerID int, jobs chan int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fmt.Printf("Worker %v started\n", workerID)

	for job := range jobs {
		webRequest(job, workerID)
	}
}

func webRequest(jobID int, workerID int) {
	_, responseBody, _ := gorequest.New().Get("http://localhost:3000/").End()
	fmt.Printf("Request %v completed by worker %v: %v\n", jobID, workerID, responseBody)
}
