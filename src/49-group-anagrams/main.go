package main

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}
	hm := make(map[[26]int][]string)
	for _, str := range strs {
		k := [26]int{}
		for _, r := range str {
			k[r-'a']++
		}
		hm[k] = append(hm[k], str)
	}
	result := [][]string{}
	for _, v := range hm {
		result = append(result, v)
	}
	return result
}
