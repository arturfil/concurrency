package waitgroups

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func WaitGroupsMain() {
	var wg sync.WaitGroup
    // wg.Add(1)

	// fmt.Println("Go waitgroup")
	// go myFunc(&wg)
    // wg.Wait()
    urls := []string {
        "https://google.com",
        "https://twitter.com",
        "https://youtube.com",
        "https://udemy.com",
    }

    /*for _, url := range urls {
        go getStatus(&wg, url)
        wg.Add(1)
    }

    wg.Wait() // wait for other waitgroups
    */

    getStatusV1(&wg, urls)
    
	fmt.Println("finished my go program")
}

func myFunc(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("Finished executing go-routine")
    wg.Done()
}

func getStatusV1(wg *sync.WaitGroup, urls []string) {
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

func getStatus(wg *sync.WaitGroup, url string) {
    defer wg.Done()

    res, err := http.Get(url) 
    if err != nil {
        fmt.Println("Error: couldn't reach url", url) 
    } else {
        fmt.Printf("%d status of %s\n", res.StatusCode, url)
    }
}
