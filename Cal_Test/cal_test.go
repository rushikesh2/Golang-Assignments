package main

import (
	"testing"
)

func TestCal(t *testing.T) {
	if calculator(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-2, 0},
		//{-3, 1},
	}

	for _, test := range tests {
		if output := calculator(test.input); output != test.expected {
			t.Error("Test failed : {} inputted, {} expected, {} received", test.input, test.expected, output)
		}

	}
}
