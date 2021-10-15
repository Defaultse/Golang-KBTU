package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan string, results chan<- string) {
	for i := range jobs {
		fmt.Println("worker", id, "started", i)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished", i)
		results <- i + "is done"
	}
}

func main() {
	cntJobs := 7
	jobsList := []string{"job1", "job2", "job3", "job4", "job5", "job6", "job7"}

	jobs := make(chan string, cntJobs)
	results := make(chan string, cntJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for _, job := range jobsList {
		jobs <- job
	}
	close(jobs)

	for i := 1; i <= cntJobs; i++ {
		<-results
	}
}
