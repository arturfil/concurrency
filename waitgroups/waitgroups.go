package concurrency

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// wait groups wait for a collection of concurrent tasks
func WaitGroupsMain() {
    var wg sync.WaitGroup

    start := time.Now()
    /*
    go loopNums(&wg, 0, 100)
    go loopNums(&wg, 100, 200)
    go loopNums(&wg, 200, 300)

    wg.Wait()
    */

    urls := []string {
        "https://google.com",
        "https://youtube.com",
        "https://udemy.com",
        "https://twitter.com",
    }

    getStatus(&wg, urls) 
    
    /*
    go func() {
        time.Sleep(time.Second * 3) 
        for i := 0; i < 100; i++ {
            fmt.Println(i)
        }
    }()
    
    go func() {
        time.Sleep(time.Second * 3) 
        for i := 100; i < 200; i++ {
            fmt.Println(i)
        }
    }() 

    go func() {
        time.Sleep(time.Second * 3) 
        for i := 200; i < 300; i++ {
            fmt.Println(i)
        }
    }()
    */

    elapsed := time.Since(start)
    fmt.Printf("Time taken %f s\n", float32(elapsed/1000000000))

} 

func loopNums(wg *sync.WaitGroup, init int, end int) {
    time.Sleep(time.Second * 3)
    for i := init; i < end; i++ {
        fmt.Println(i)
    }
    wg.Done()
}

func getStatus(wg *sync.WaitGroup, urls []string) {
    for _, url := range urls {
        wg.Add(1)
        go func(url string) {
            res, err := http.Get(url) 

            if err != nil {
                fmt.Printf("[Error: %s x -> %s]\n", err, url)
            } else {
                fmt.Printf("[%d %s]\n", res.StatusCode, url)
            }
            wg.Done()
        }(url)
    } 
    wg.Wait()
}


