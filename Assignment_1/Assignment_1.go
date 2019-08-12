package main

import (
    "encoding/json"
    "fmt"
   	"io/ioutil"
   	"os"
)

type company struct {
	Stuff struct {
		Onetype []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"onetype"`
		Othertype struct {
			ID      int    `json:"id"`
			Company string `json:"company"`
		} `json:"othertype"`
	} `json:"stuff"`
	Otherstuff struct {
		Thing [][]int `json:"thing"`
	} `json:"otherstuff"`
}

func UnmanJson(file string) {
     var com company
     datafile, err := os.Open(file)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println("File Sucessfully opened :")

        defer datafile.Close()

        //read json file
        temp , _ := ioutil.ReadAll(datafile)

        //unmarsh json data
         er := json.Unmarshal(temp, &com)
         if er != nil {
             fmt.Println("There was an error:", err)
         }
           fmt.Println(com.Stuff)
             fmt.Println(com)

}

func main() {
    UnmanJson("user.json")


 /*   //open json file
    datafile, err := os.Open("user.json")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("File Sucessfully opened :")

    defer datafile.Close()

    //read json file
    temp , _ := ioutil.ReadAll(datafile)

    //unmarsh json data
     er := json.Unmarshal(temp, &com)
     if er != nil {
         fmt.Println("There was an error:", err)
     }
*/

}