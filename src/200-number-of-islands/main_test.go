package main

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	tests := []struct {
		grid   [][]byte
		expect int
	}{
		{
			grid: [][]byte{
				[]byte("11110"),
				[]byte("11010"),
				[]byte("11000"),
				[]byte("00000"),
			},
			expect: 1,
		},
		{
			grid: [][]byte{
				[]byte("11000"),
				[]byte("11000"),
				[]byte("00100"),
				[]byte("00011"),
			},
			expect: 3,
		},
	}

	for _, test := range tests {
		actual := numIslands(test.grid)
		if actual != test.expect {
			t.Errorf("For grid %v, expected %d islands, but got %d", test.grid, test.expect, actual)
		}
	}
}
