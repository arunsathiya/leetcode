package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		kValue   int
		expected []int
	}{
		{
			name:     "test case from leetcode",
			input:    []int{1, 1, 1, 2, 2, 3},
			kValue:   2,
			expected: []int{1, 2},
		},
		{
			name:     "second test case from leetcode",
			input:    []int{1},
			kValue:   1,
			expected: []int{1},
		},
		{
			name:     "third test case from leetcode",
			input:    []int{3, 0, 1, 0},
			kValue:   1,
			expected: []int{0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := topKFrequent(test.input, test.kValue)
			if !reflect.DeepEqual(output, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, output)
			}
		})
	}
}
