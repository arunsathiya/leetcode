package main

import (
	"testing"
)

func TestTrap(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := trap(tt.height)
			if got != tt.want {
				t.Errorf("trap(%v) = %v, want %v", tt.height, got, tt.want)
			}
		})
	}
}
