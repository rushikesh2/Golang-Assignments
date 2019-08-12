package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func main() {

	fmt.Println("main() started", time.Since(start))
	chan1 := make(chan string, 2)
	chan2 := make(chan string, 2)

	chan2 <- "Value 3"
	chan1 <- "Value 2"
	chan2 <- "Value 1"
	chan1 <- "Value 2"
	//chan2 <- "Value 8"


	select {
	case res := <-chan1:
		fmt.Println("Response from chan1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from chan2", res, time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}
