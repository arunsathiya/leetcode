package main

import "testing"

func TestRomanToInt(t *testing.T) {
	cases := []struct {
		s        string
		expected int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, c := range cases {
		got := romanToInt(c.s)
		if got != c.expected {
			t.Errorf("romanToInt(%q) == %d, want %d", c.s, got, c.expected)
		}
	}
}
