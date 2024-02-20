package main

import "strings"

func romanToInt(s string) int {
	characters := strings.Split(s, "")
	count := 0
	skip := false
	for index, char := range characters {
		if skip {
			skip = false
			continue
		}
		switch char {
		case "I":
			if index+1 < len(characters) && characters[index+1] == "V" {
				count += 4
				skip = true
				continue
			}
			if index+1 < len(characters) && characters[index+1] == "X" {
				count += 9
				skip = true
				continue
			}
			count += 1
		case "V":
			count += 5
		case "X":
			if index+1 < len(characters) && characters[index+1] == "L" {
				count += 40
				skip = true
				continue
			}
			if index+1 < len(characters) && characters[index+1] == "C" {
				count += 90
				skip = true
				continue
			}
			count += 10
		case "L":
			count += 50
		case "C":
			if index+1 < len(characters) && characters[index+1] == "D" {
				count += 400
				skip = true
				continue
			}
			if index+1 < len(characters) && characters[index+1] == "M" {
				count += 900
				skip = true
				continue
			}
			count += 100
		case "D":
			count += 500
		case "M":
			count += 1000
		}
	}
	return count
}
