package main

import (
	"reflect"
	"testing"
)

func TestSortPeople(t *testing.T) {
	tests := []struct {
		name     string
		names    []string
		heights  []int
		expected []string
	}{
		{
			name:     "first test case from leetcode",
			names:    []string{"Mary", "John", "Emma"},
			heights:  []int{180, 165, 170},
			expected: []string{"Mary", "Emma", "John"},
		},
		{
			name:     "second test case from leetcode",
			names:    []string{"Alice", "Bob", "Bob"},
			heights:  []int{155, 185, 150},
			expected: []string{"Bob", "Alice", "Bob"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sortPeople(test.names, test.heights)
			if !reflect.DeepEqual(output, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, output)
			}
		})
	}
}
