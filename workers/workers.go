package workers

import (
	"fmt"
	"time"
)

func WorkersMain() {
	const numbJobs = 10
	start := time.Now()
    completed := []int{}

	jobsChan := make(chan int, numbJobs) // buffered 
	completedJobsChan := make(chan int, numbJobs) 

	for i := 1; i <= 3; i++ { // numer of workers => 2 | time take => 1 => 10 / (workers * time) = 10 / 2 => 5
		go worker(i, jobsChan, completedJobsChan)
	}

    // blocking code, loads channels numbers, number of jobs = 10
	for j := 1; j <= numbJobs; j++ {
		jobsChan <- j 	
    }

	close(jobsChan) // once loaded the jobs

    // frees up the channels info & append the result to completed slice
	for k := 1; k <= numbJobs; k++ {
        proccesed := <- completedJobsChan
        completed = append(completed, proccesed)
	}

	end := time.Since(start)
    fmt.Println("completedJobsChan", completed)
	fmt.Println("It has taken: ", end)
}

// only read, only send
func worker(id int, jobsChan <-chan int, completedJobsChan chan<- int) {
	for job := range jobsChan {
        time.Sleep(time.Second * 1) // should take 1 second per worker
		fmt.Println("worker\t", id, "\t[started]\t job", job, "with", len(jobsChan), "jobs left to process")
		fmt.Println("worker\t", id, "\t[finished]\t job", job)
        // val := job * job
        completedJobsChan <- job * job // squared
	}
}
