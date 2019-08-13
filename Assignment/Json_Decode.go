package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	datafile, err := os.Open("user.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File sucessfully opened")

	defer datafile.Close()

	//read json file
	temp, _ := ioutil.ReadAll(datafile)

	var result map[string]interface{}
	json.Unmarshal([]byte(temp), &result)

	fmt.Println(result["stuff"])

}
