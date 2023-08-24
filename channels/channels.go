package channels

import (
	"fmt"
	"time"
)

func ChannelsMain() {
    // values := make(chan int)
    values := make(chan string, 2) // this 2, is a capacity for the queue of no. of channles proc
    defer close(values)

    go sendString(values)
    go sendString(values)

    value := <- values
    fmt.Println(value)
}

func sendValue(c chan int) {
    c <- 8 
}

func sendString(c chan string) {
    fmt.Println("[START]")
    time.Sleep(1 * time.Second)
    c <- "Hello world"
    fmt.Println("[FINISH]")
}
