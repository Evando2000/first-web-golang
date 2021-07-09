package main

import "testing"

func TestCalc(t *testing.T) {
	if calc(2) != 4 {
		t.Error("Expected 2 + 2 = 4")
	}
}
