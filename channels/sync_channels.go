package channels

import (
	"fmt"
	"time"
)

var result = 0
var value = 97

func SycnMain() {
    goChan := make(chan int)
    mainChan := make(chan string)

    go calculateSquare(value, goChan)
    go reportResult(goChan, mainChan)

    <- mainChan
}

func calculateSquare(value int, goChan chan int) {
    fmt.Println("Calculating for 3 seconds...")
    time.Sleep(time.Second * 3)
    result = value * value
    goChan <- result
}

func reportResult(goChan chan int, mainChan chan string) {
    time.Sleep(time.Second * 1)
    fmt.Println("The result of ", value, "squared is", <- goChan)
    mainChan <- "Process done..." 
}
