package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "os"
)

func main() {

    datafile, err := os.Open("user1.json")
    if err != nil {
           fmt.Println(err)
       }
     fmt.Println("File sucessfully opened")

    defer datafile.Close()

    // Declared an empty interface of type Array
    var result []map[string]interface{}

    json.Unmarshal([]byte(empArray), &result)

    for key, result := range result {
        fmt.Println("Reading Value for Key :", key)
        fmt.Println("Id :", result["id"],
        			"- Name :", result["name"],
        			"- Department :", result["department"],
        			"- Designation :", result["designation"])
    }