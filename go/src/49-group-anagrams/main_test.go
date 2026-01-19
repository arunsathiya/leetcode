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
			
			// Sort strings within each group
			for _, g := range got {
				sort.Strings(g)
			}
			for _, w := range tt.want {
				sort.Strings(w)
			}
			
			// Sort groups themselves for consistent comparison
			sort.Slice(got, func(i, j int) bool {
				return got[i][0] < got[j][0]
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i][0] < tt.want[j][0]
			})
			
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams(%v) = %v, want %v", tt.strs, got, tt.want)
			}
		})
	}
}
