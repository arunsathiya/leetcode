package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		target   int
		expected []int
	}{
		{
			name:     "test case from leetcode",
			input:    []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "test case from leetcode",
			input:    []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := twoSum(test.input, test.target)
			if !reflect.DeepEqual(test.expected, output) {
				t.Errorf("Expected %v, got %v", test.expected, output)
			}
		})
	}
}
