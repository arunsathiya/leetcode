package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:     "test case from leetcode",
			input:    []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := threeSum(test.input)
			if !reflect.DeepEqual(output, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, output)
			}
		})
	}
}
