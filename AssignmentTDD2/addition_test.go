package assignmenttdd2

import (
	"testing"
)

func TestAddition(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{1, 4},
		{0, 3},
		{-1, 2},
		{-3, 0},
	}
	for _, test := range tests {
		if result := Addition(test.input); result != test.expected {
			t.Error("Test failed : {} inputted, {} expected, {} received", test.input, test.expected, result)
		}
	}
}
