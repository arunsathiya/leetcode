package main

import "testing"

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		word1 string
		word2 string
		want  string
	}{
		{"abc", "pqr", "apbqcr"},
		{"ab", "pqrs", "apbqrs"},
		{"abcd", "pq", "apbqcd"},
	}

	for _, tt := range tests {
		got := mergeAlternately(tt.word1, tt.word2)
		if got != tt.want {
			t.Errorf("mergeAlternately(%q, %q) = %q, want %q", tt.word1, tt.word2, got, tt.want)
		}
	}
}
