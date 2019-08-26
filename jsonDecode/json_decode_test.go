package jsondecode

import "testing"

var input1 = `{"name": "xyz"}`

var input2 = `{
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

var input3 = `hello`

var testcases = []string{input1, input2, input3}

func TestDecode(t *testing.T) {
	for _, val := range testcases {
		DecodeJSON(val)
		//fmt.Println(val)
	}

}
