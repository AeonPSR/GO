package main

import (
	"testing"
)

//The name convention itself is what GO uses to get what should test which test file.

func TestSomme(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{ //Numbers to add up, and the result that shoulds come up after the }.
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{0, 0, 0, 0, 0}, 0},
		{[]int{-1, -2, -3, -4, -5}, -15},
		{[]int{10, 20, 30, 40, 50}, 150},
		{[]int{}, 0},  // Empty
	}

	for _, test := range tests {
		result := somme(test.input)
		if result != test.expected {
			t.Errorf("somme(%v) = %d; attendu %d", test.input, result, test.expected)
		}
	}
}
