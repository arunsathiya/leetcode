package testutil

import (
	"encoding/json"
)

// Parse2DIntArray converts LeetCode format string to [][]int
// Example: "[[1,2],[3,4]]" -> [][]int{{1,2},{3,4}}
func Parse2DIntArray(jsonStr string) [][]int {
	var result [][]int
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		panic("Failed to parse 2D int array: " + err.Error())
	}
	return result
}
