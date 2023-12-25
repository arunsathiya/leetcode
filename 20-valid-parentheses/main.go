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

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func isValid(s string) bool {
	var stack Stack
	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack.Push(string(char))
		case ')', ']', '}':
			if stack.IsEmpty() {
				return false
			}
			top := stack.Pop()
			if (char == '(' && top != ")") || (char == '[' && top != "]") || (char == '{' && top != "}") {
				return false
			}
		}
	}
	return stack.IsEmpty()
}
