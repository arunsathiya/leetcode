package main

func numberOfWays(s string) int64 {
	totalZeroes := 0
	totalOnes := 0
	countZeroesSoFar := 0
	countOnesSoFar := 0
	valid101Combos := 0
	valid010Combos := 0
	for _, char := range s {
		if char == '0' {
			totalZeroes++
		} else {
			totalOnes++
		}
	}
	for _, char := range s {
		if char == '0' {
			countZeroesSoFar++
			valid101Combos += countOnesSoFar * (totalOnes - countOnesSoFar)
		} else {
			countOnesSoFar++
			valid010Combos += countZeroesSoFar * (totalZeroes - countZeroesSoFar)
		}
	}
	return int64(valid010Combos + valid101Combos)
}
