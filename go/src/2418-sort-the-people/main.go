package main

import "sort"

type Pair struct {
	Name   string
	Height int
}

type ByHeight []Pair

func (a ByHeight) Len() int           { return len(a) }
func (a ByHeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByHeight) Less(i, j int) bool { return a[i].Height < a[j].Height }

func sortPeople(names []string, heights []int) []string {
	pairs := make(ByHeight, len(names))
	results := make([]string, 0)
	for i, name := range names {
		pairs[i] = Pair{name, heights[i]}
	}
	sort.Sort(pairs)
	for i := len(pairs) - 1; i >= 0; i-- {
		results = append(results, pairs[i].Name)
	}
	return results
}
