package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		width, height, length, mass float64
		expected                    string
	}{
		// Standard packages
		{90, 90, 90, 19, "STANDARD"},
		// Bulky packages
		{150, 150, 150, 19, "SPECIAL"},
		{100, 100, 100, 19, "SPECIAL"},
		// Heavy packages
		{90, 90, 90, 20, "SPECIAL"},
	}

	for _, test := range tests {
		result := solve(test.width, test.height, test.length, test.mass)
		if result != test.expected {
			t.Errorf("solve(%f, %f, %f, %f) = %s; expected %s", test.width, test.height, test.length, test.mass, result, test.expected)
		}
	}
}
