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
    // bufferedChannels()
    // channelSelect()
}


func unbufferendChannels() {
    c := make(chan int)

    start := time.Now()

    for i := 0; i < 3; i++ {
        go func() {
            for val := range c {
                time.Sleep(time.Millisecond * 500) // 10 / 4  -> 2.5 * 2 ~ 5s 
                fmt.Println("value -> ", val)
            }
        }()
    }

    for i := 1; i <= 10; i++ {
        c <- i
    }
    close(c)

    end := time.Since(start)
    fmt.Println("It has taken: ", end)
}


func nonSeqUnbuffered() {
    c := make(chan int)

    start := time.Now()

    for i := 0; i < 3; i++ { // number of workers -> should take only two seconds
        wg.Add(1)
        go func() {
            defer wg.Done()
            for val := range c {
                time.Sleep(time.Millisecond * 500) // 10 / 4  -> 2.5 * 2 ~ 5s 
                fmt.Println("value -> ", val)
            } 
        }()
    }
        
    for i := 1; i <= 10; i++ {
        c <- i
    }

    close(c)
    wg.Wait()

    end := time.Since(start)

    fmt.Println("It has taken: ", end)
}

// use case -> order matters & comunication between processes
func bufferedChannels() {
    // no need for go routines
    // after that, it behaves exactly the same as an unbuffered channel
    c := make(chan string, 4)

    c <- "Hello"
    c <- "My Name is"
    c <- "Arturo"
    c <- "I like to code and play soccer"

    fmt.Println(<- c)
    fmt.Println(<- c)
    fmt.Println(<- c)
    fmt.Println(<- c)
}


func channelSelect() {
    c1 := make(chan string)
    c2 := make(chan string)
    c3 := make(chan string)

    go func() {
        for {
            time.Sleep(time.Millisecond * 500) 
            c1 <- "Every 500ms"
        }
    }()

    go func() {
        for {
            time.Sleep(time.Second * 2)
            c2 <- "Every two seconds"
        }
    }()

    go func() {
        for {
            time.Sleep(time.Second * 8)
            c3 <- "Process finished..."
        }
    }()

    for {
        select {
            case ms1 := <- c1:
                fmt.Println(ms1)
            case ms2 := <- c2: 
                fmt.Println(ms2)
            case ms3 := <- c3:
                fmt.Println(ms3)
                os.Exit(0)
        }
    }  
}

