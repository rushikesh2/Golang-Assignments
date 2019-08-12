package main

import "fmt"

func main() {
    fmt.Println("main started")
    c := make(chan string)

    //Anonymous goroutine
    go func(c chan string) {
        fmt.Println("Hello"+ <- c +  "!")
    }(c)

    //Blocked
    c <- "Rushi"

    fmt.Println("main started")
}