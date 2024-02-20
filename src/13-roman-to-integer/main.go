package main

import "strings"

func romanToInt(s string) int {
	characters := strings.Split(s, "")
	count := 0
	for index, char := range characters {
		switch char {
		case "I":
			if index+1 < len(characters) && characters[index+1] == "V" {
				count += 4
			} else if index+1 < len(characters) && characters[index+1] == "X" {
				count += 9
				break
			} else {
				count += 1
			}
		case "V":
			count += 5
		case "X":
			if index+1 < len(characters) && characters[index+1] == "L" {
				count += 40
			} else if index+1 < len(characters) && characters[index+1] == "C" {
				count += 90
				break
			} else {
				count += 10
			}
		case "L":
			count += 50
		case "C":
			if index+1 < len(characters) && characters[index+1] == "D" {
				count += 400
			} else if index+1 < len(characters) && characters[index+1] == "M" {
				count += 900
				break
			} else {
				count += 100
			}
		case "D":
			count += 500
		case "M":
			count += 1000
		}
	}
	return count
}
