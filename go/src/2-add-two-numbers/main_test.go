package main

import (
	"testing"
)

func TestaddTwoNumbers(t *testing.T) {
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
			// got := addTwoNumbers(...)
			// if got != tt.want {
			//     t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func BenchmarkaddTwoNumbers(b *testing.B) {
	// TODO: add benchmark inputs
	for i := 0; i < b.N; i++ {
		// addTwoNumbers(...)
	}
}
