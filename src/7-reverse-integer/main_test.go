package main

import (
	"testing"
)

func Testreverse(t *testing.T) {
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
			// got := reverse(...)
			// if got != tt.want {
			//     t.Errorf("reverse() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func Benchmarkreverse(b *testing.B) {
	// TODO: add benchmark inputs
	for i := 0; i < b.N; i++ {
		// reverse(...)
	}
}
