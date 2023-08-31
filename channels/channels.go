package channels

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func ChannelsMain() {
    c := make(chan int, 10)

    go func() {
        var wg sync.WaitGroup
        for i := 0; i < 10; i++ {
            wg.Add(1)
            go func(i int) {
                defer wg.Done()
                c <- i 
                time.Sleep(time.Second * 5)
            }(i)
        }
        wg.Wait()
        close(c)
    }()

    for val := range c {
        fmt.Println(val)
    }

    // channelSelect()

    // values := make(chan int)
    /*
    values := make(chan string, 2) // this 2, is a capacity for the queue of no. of channles proc
    defer close(values)

    go sendString(values)
    go sendString(values)

    value := <- values
    fmt.Println(value)
    */
    /*
    c := make(chan string)
    go count("sheep", c)

    for msg := range c {
        fmt.Println(msg)
    }
    */
}

func count(item string, c chan string) {
    for i := 0; i < 5; i++ {
        c <- item 
        time.Sleep(time.Microsecond * 500)
    }
    close(c)
}

func sendValue(c chan int) {
    c <- 8 
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

func sendString(c chan string) {
    fmt.Println("[START]")
    time.Sleep(1 * time.Second)
    c <- "Hello world"
    fmt.Println("[FINISH]")
}
