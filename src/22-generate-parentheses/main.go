package main

func backtrack(result *[]string, path string, open, close, max int) {
	if len(path) == max*2 {
		*result = append(*result, path)
		return
	}
	if open < max {
		backtrack(result, path+"(", open+1, close, max)
	}
	if close < open {
		backtrack(result, path+")", open, close+1, max)
	}
}

func generateParenthesis(n int) []string {
	var result []string
	backtrack(&result, "", 0, 0, n)
	return result
}
