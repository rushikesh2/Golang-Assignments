package assignments

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestMakeunmarshal(t *testing.T) {

	datafile, err := os.Open("user.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File Sucessfully opened :")

	defer datafile.Close()

	temp, _ := ioutil.ReadAll(datafile)
	Makeunmarshal(temp)
	if err != nil {
		t.Error()
	}
}
