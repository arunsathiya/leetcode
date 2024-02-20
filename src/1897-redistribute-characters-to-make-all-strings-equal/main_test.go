package main

import "testing"

func TestMakeEqual(t *testing.T) {
	tests := []struct {
		words []string
		want  bool
	}{
		{[]string{"abc", "aabc", "bc"}, true},
		{[]string{"ab", "a"}, false},
		{[]string{"b", "a"}, false},
	}
	for _, tt := range tests {
		got := makeEqual(tt.words)
		if got != tt.want {
			t.Errorf("makeEqual(%v) = %v, want %v", tt.words, got, tt.want)
		}
	}
}
