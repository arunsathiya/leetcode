package main

func numberOfWays(s string) int64 {
	count0, count1 := 0, 0
	count01, count10 := 0, 0
	ways := 0
	for _, char := range s {
		if char == '0' {
			ways += count01
			count10 += count1
			count0++
		} else {
			ways += count10
			count01 += count0
			count1++
		}
	}
	return int64(ways)
}
