package main

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}
	anagrams := make(map[string][]string)
	for _, str := range strs {
		sortedString := sortString(str)
		anagrams[sortedString] = append(anagrams[sortedString], str)
	}
	result := [][]string{}
	for _, v := range anagrams {
		result = append(result, v)
	}
	return result
}
