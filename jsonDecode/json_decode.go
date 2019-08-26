package jsondecode

import (
	"encoding/json"
	"fmt"
)

var jsonstring = `{
"stuff": {
"onetype": [
{"id":1,"name":"John Doe"},
{"id":2,"name":"Don Joeh"}	
],
"othertype": {"id":2,"company":"ACME"}
},
"otherstuff": {
"thing": [[1,42],[2,2]]
}
}`

//DecodeJSON function is used to Unmarshal the JSON input.
func DecodeJSON(jsonstring string) {
	stuffdata := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsonstring), &stuffdata)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	ParseMap(stuffdata)
}

//ParseMap function is used for parsing map.
func ParseMap(aMap map[string]interface{}) {
	for key, val := range aMap {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println(key)
			ParseMap(val.(map[string]interface{}))
		case []interface{}:
			fmt.Println(key)
			ParseArray(val.([]interface{}))
		default:
			fmt.Println(key, ":", concreteVal)
		}
	}
}

//ParseArray function is used to parse an array.
func ParseArray(anArray []interface{}) {
	for i, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println("Index:", i)
			ParseMap(val.(map[string]interface{}))
		case []interface{}:
			fmt.Println("Index:", i)
			ParseArray(val.([]interface{}))
		default:
			fmt.Println("Index", i, ":", concreteVal)

		}
	}
}
