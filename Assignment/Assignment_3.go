package main

import (
	"encoding/json"
	"fmt"
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

	temp, _ := ioutil.ReadAll(datafile)
	// Declared an empty interface of type Array
	var result []map[string]interface{}

	json.Unmarshal([]byte(temp), &result)

	for key, result := range result {
		fmt.Println("Reading Value for Key :", key)
		fmt.Println("Id :", result["id"],
			"- Name :", result["name"],
			"- Department :", result["department"],
			"- Designation :", result["designation"])

	}
}
