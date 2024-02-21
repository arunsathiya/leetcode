package main

func partitionString(s string) int {
	optimal := make([]string, 0)
	for len(s) > 0 {
		instance := ""
		for s[0] != instance[len(instance)-1] {
			instance += string(s[0])
			s = s[1:]
		}
		optimal = append(optimal, instance)
	}
	return len(optimal)
}
