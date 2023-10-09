package channels

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func ChannelsMain() {
    // unbufferendChannels()
    // nonSeqUnbuffered()
    bufferedChannels()
    // bufferedChannelsWithWorkers()
    // channelSelect()
    // workers.WorkersMain()
}


func unbufferendChannels() {
    c := make(chan int)
    processed_vals := []int{}
    start := time.Now()

    // for i := 0; i < 3; i++ {
    // for w := 1; w<=3; w++ { // 3 workers
        go func() {
            for val := range c {
                fmt.Println("value -> ", val) // flip this to show error
                val *= val // squaring
                val += 1
                processed_vals = append(processed_vals, val)
                time.Sleep(time.Millisecond * 500) // 10 / 4  -> 2.5 * 2 ~ 5s 
            }
            close(c)
        }()

    // }
        
    // blocking code
    for i := 1; i <= 10; i++ {
        c <- i
    }
    // close(c)

    fmt.Println("Processed vals", processed_vals)

    end := time.Since(start)
    fmt.Println("It has taken: ", end)
}


func nonSeqUnbuffered() {
    c := make(chan int)
    processed_vals := []int{}

    start := time.Now()

    for worker := 0; worker < 3; worker++ { // number of workers -> should take only two seconds

        wg.Add(1)
        go func() {
            defer wg.Done()
            for val := range c {
                time.Sleep(time.Millisecond * 500) // 10 / 4  -> 2.5 * 2 ~ 5s 
                fmt.Println("value -> ", val)
                processed_vals = append(processed_vals, val)
            } 
        }()

    }
        
    // blocking code 
    for i := 1; i <= 10; i++ {
        c <- i
    }

    close(c)
    wg.Wait()

    fmt.Println("Processed vals", processed_vals)

    end := time.Since(start)
    fmt.Println("It has taken: ", end)
}

// use case -> order matters & comunication between processes
func bufferedChannels() {
    // no need for go routines. After that, it behaves exactly the same as an unbuffered channel
    // are blocking code with a queue of length of the channel length

    // c := make(chan string, 2) // creating buffer length

    // c <- "Hello"
    // c <- "My Name is Arturo"

    // fmt.Println(<- c)
    // fmt.Println(<- c)

    // c <- "Im showing concurrency"
    // fmt.Println(<- c)

    
    c := make(chan int, 2)
    start := time.Now()
    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // work done
            // blocking code
            for val := range c {
                fmt.Println("Read\t", val, "\tfrom channel")
                time.Sleep(time.Millisecond * 2000)
            }
                    
        }()
    }
    
    // blockin code
    for i := 0; i < 4; i++ { // tot of 4 rounds
        c <- i
        fmt.Println("Wrote\t", i, "\tto channel")
    }

    close(c)
    wg.Wait()

    end := time.Since(start)
    fmt.Println("Time taken", end)
}


func bufferedChannelsWithWorkers() {
    c := make(chan int, 2)
    start := time.Now()

    for i := 0; i < 2; i++ { // number of workers
        wg.Add(1)
        go func() {
            defer wg.Done()
            for val := range c {
                fmt.Println("Read\t", val, "\tfrom channel")
                time.Sleep(time.Millisecond * 2000)
            }
        }()
    }

    // blocking code
    for i := 0; i < 4; i++ { // tot of 4 rounds
        c <- i
        fmt.Println("Wrote\t", i, "\tto channel")
    }

    close(c)
    wg.Wait()

    end := time.Since(start)
    fmt.Println("Time taken", end)
}


func channelSelect() {
    c1 := make(chan string)
    c2 := make(chan string)
    c3 := make(chan string)

    start := time.Now()


    go func() {
        for {
            time.Sleep(time.Millisecond * 500) 
            c1 <- "Every 500ms"
        }
    }()

    go func() {
        for {
            time.Sleep(time.Millisecond * 2000)
            // time.Sleep(time.Second * 2) // w/ select
            c2 <- "Every two seconds"
        }
    }()

    go func() {
        for {
            time.Sleep(time.Second * 8)
            // time.Sleep(time.Second * 8) // w/ select example
            c3 <- "Process finished..."
            // end := time.Since(start)
            // fmt.Println("It has taken: ", end)            
            os.Exit(0) // example without select
        }
    }()

    // for {
    //     fmt.Println(<- c1)
    //     fmt.Println(<- c2)
    //     fmt.Println(<- c3)
    // }

    for {
        select {
            case ms1 := <- c1:
                fmt.Println(ms1)
            case ms2 := <- c2: 
                fmt.Println(ms2)
            case ms3 := <- c3:
                fmt.Println(ms3)
                end := time.Since(start)
                fmt.Println("It has taken: ", end)
                os.Exit(0)
        }
    }  

    
}

