package main

import "testing"

func TestNumberOfWays(t *testing.T) {
	tests := []struct {
		s        string
		expected int64
	}{
		{"001101", 6},
		{"11100", 0},
	}

	for _, test := range tests {
		if got := numberOfWays(test.s); got != test.expected {
			t.Errorf("numberOfWays(%q) = %d; want %d", test.s, got, test.expected)
		}
	}
}
