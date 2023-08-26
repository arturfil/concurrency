package workers

import (
	"fmt"
	"time"
)

func WorkersMain() {
   const numbJobs = 10 

   jobsChan := make(chan int, numbJobs) // buffered
   completedJobsChan := make(chan int, numbJobs)

   for i := 1; i <= 3; i++ {
       go worker(i, jobsChan, completedJobsChan)
   }

   for j := 1; j <= numbJobs ; j++ {
        jobsChan <- j // blocking code, loads channels numbers
   }

   close(jobsChan)

   for a := 1; a <= numbJobs; a++ {
       <- completedJobsChan
   }
}

func worker(id int, jobsChan <- chan int, completedJobsChan chan <- int) {
    for j := range jobsChan {
        fmt.Println("worker", id, "started jobs", j, "with", len(jobsChan), "jobs left to process") 
        time.Sleep(time.Second * 2)
        fmt.Println("worker", id, "finished job", j)
        completedJobsChan <- j // load finishes job numbers
    }
}

