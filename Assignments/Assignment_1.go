package assignments

import (
	"encoding/json"
	"fmt"
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

//Makeunmarshal This function can decode the JSON string
func Makeunmarshal(file []byte) {
	var com company
	err := json.Unmarshal(file, &com)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	fmt.Printf("%+v\n", com)
}
