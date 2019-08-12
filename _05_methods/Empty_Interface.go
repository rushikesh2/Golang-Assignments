package main

import "fmt"

type MyString string

type MyInt int
func explain(i interface{}) {
    fmt.Printf("value given to explain function is of type '%T' with value is %v\n", i, i)
}
func main() {

    ms := MyString("Hello World")
    a := MyInt(5)
    explain(ms)
    explain(a)
}