package main

import (
	"reflect"
	"testing"
)

func TestSequentialDigits(t *testing.T) {
	testCases := []struct {
		name string
		low  int
		high int
		want []int
	}{
		{
			name: "Test Case 1: Range 100 to 300",
			low:  100,
			high: 300,
			want: []int{123, 234},
		},
		{
			name: "Test Case 2: Range 1000 to 13000",
			low:  1000,
			high: 13000,
			want: []int{1234, 2345, 3456, 4567, 5678, 6789, 12345},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := sequentialDigits(tc.low, tc.high)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("sequentialDigits(%d, %d) = %v, want %v", tc.low, tc.high, got, tc.want)
			}
		})
	}
}
