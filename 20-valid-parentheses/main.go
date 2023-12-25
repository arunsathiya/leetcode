package main

func isValid(s string) bool {
	for i := range s {
		if string(s[i]) == "(" && string(s[len(s)-1]) == ")" {
			return true
		} else if string(s[i]) == "[" && string(s[len(s)-1]) == "]" {
			return true
		} else if string(s[i]) == "{" && string(s[len(s)-1]) == "}" {
			return true
		}
	}
	return false
}
