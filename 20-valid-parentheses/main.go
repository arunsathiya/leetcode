package main

type Stack []string

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() string {
	if len(*s) == 0 {
		return ""
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}

func (s *Stack) Peek() string {
	if len(*s) == 0 {
		return ""
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

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
