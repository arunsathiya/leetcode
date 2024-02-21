package main

func isSubsequence(s string, t string) bool {
	sloop := 0
	for tloop := 0; tloop < len(t) && sloop < len(s); tloop++ {
		if s[sloop] == t[tloop] {
			sloop++
		}
	}
	return sloop == len(s)
}
