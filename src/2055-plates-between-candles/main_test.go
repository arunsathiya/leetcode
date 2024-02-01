package main

import (
	"reflect"
	"testing"
)

func TestPlatesBetweenCandles(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		queries  [][]int
		expected []int
	}{
		{
			name:     "first test case from leetcode",
			s:        "**|**|***|",
			queries:  [][]int{{2, 5}, {5, 9}},
			expected: []int{2, 3},
		},
		{
			name:     "second test case from leetcode",
			s:        "***|**|*****|**||**|*",
			queries:  [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}},
			expected: []int{9, 0, 0, 0, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := platesBetweenCandles(test.s, test.queries)
			if !reflect.DeepEqual(output, test.expected) {
				t.Errorf("Expected %d, got %d", test.expected, output)
			}
		})
	}
}
