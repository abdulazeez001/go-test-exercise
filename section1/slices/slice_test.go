package slices

import (
	"testing"
)

func TestSumOfAllEvenNum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Nil slice", nil, 0},
		{"Empty slice", []int{}, 0},
		{"All odd numbers", []int{1, 3, 5, 7}, 0},
		{"All even numbers", []int{2, 4, 6, 8}, 20},
		{"Mixed numbers", []int{1, 2, 3, 4, 5, 6}, 12},
		{"Single even number", []int{2}, 2},
		{"Single odd number", []int{1}, 0},
		{"Large numbers", []int{1000000000, 1000000001, 1000000002, 1000000003}, 2000000002},
		{"Negative numbers", []int{-2, -3, -4, -5}, -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumOfAllEvenNum(tt.input)
			if result != tt.expected {
				t.Errorf("Expected sum of even numbers to be %d, but got %d", tt.expected, result)
			}
		})
	}
}

func TestIsValEven(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Even number", 4, true},
		{"Odd number", 5, false},
		{"Zero", 0, true},
		{"Negative even number", -2, true},
		{"Negative odd number", -3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValEven(tt.input)
			if result != tt.expected {
				t.Errorf("Expected isValEven(%d) to be %v, but got %v", tt.input, tt.expected, result)
			}
		})
	}
}
