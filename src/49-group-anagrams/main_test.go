package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs []string
		want [][]string
	}{
		{
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			strs: []string{""},
			want: [][]string{{""}},
		},
		{
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		t.Run("groupAnagrams", func(t *testing.T) {
			got := groupAnagrams(tt.strs)
			for _, g := range got {
				sort.Strings(g) // Sort to avoid mismatch due to different order
			}
			for _, w := range tt.want {
				sort.Strings(w) // Sort to ensure the comparison works correctly
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams(%v) = %v, want %v", tt.strs, got, tt.want)
			}
		})
	}
}
