package main

import (
    "fmt"
    "sync"
)

//i = 0
var i int

// goroutine increment global variable i
func worker(wg *sync.WaitGroup) {
    i = i + 1
    wg.Done()
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i <=1000; i++{
        wg.Add(1)
        go worker(&wg)
    }

	// wait until all 1000 gorutines are done
    wg.Wait()
    fmt.Println("value of i after 1000 operation is :", i)
}