package main

import (
	"testing"
)

func TestareaOfMaxDiagonal(t *testing.T) {
	tests := []struct {
		name string
		// TODO: add test fields
		want interface{}
	}{
		{
			name: "Example 1",
			// TODO: add test case
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: implement test
			// got := areaOfMaxDiagonal(...)
			// if got != tt.want {
			//     t.Errorf("areaOfMaxDiagonal() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func BenchmarkareaOfMaxDiagonal(b *testing.B) {
	// TODO: add benchmark inputs
	for i := 0; i < b.N; i++ {
		// areaOfMaxDiagonal(...)
	}
}
