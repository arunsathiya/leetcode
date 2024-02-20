package main

var hm = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var value, lastValue, currentValue int
	for i := len(s) - 1; i >= 0; i-- {
		currentValue = hm[s[i]]
		if currentValue < lastValue {
			value -= currentValue
		} else {
			value += currentValue
		}
		lastValue = currentValue
	}
	return value
}
