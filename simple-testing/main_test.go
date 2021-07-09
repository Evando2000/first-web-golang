package main

import "testing"

func TestCalc(t *testing.T) {
	if calc(2) != 4 {
		t.Error("Expected 2 + 2 = 4")
	}
}

func TestTableCalc(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{99999, 100001},
	}
	for _, test := range tests {
		if output := calc(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
