package main

type Stack []rune

func (s *Stack) Push(r rune) {
	*s = append(*s, r)
}

func (s *Stack) Pop() rune {
	if len(*s) == 0 {
		return 0
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}

func (s *Stack) Peek() rune {
	if len(*s) == 0 {
		return 0
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
			stack.Push(char)
		case ')', ']', '}':
			if stack.IsEmpty() {
				return false
			}
			top := stack.Pop()
			if (char == '(' && top != ')') || (char == '[' && top != ']') || (char == '{' && top != '}') {
				return false
			}
		}
	}
	return stack.IsEmpty()
}
