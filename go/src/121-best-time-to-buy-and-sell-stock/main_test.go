package main

import "testing"

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "Example 1",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "Example 2",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := maxProfit(tc.prices)
			if got != tc.want {
				t.Errorf("maxProfit(%v) = %v, want %v", tc.prices, got, tc.want)
			}
		})
	}
}
