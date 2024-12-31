package main

func romanToInt(s string) int {
	hash := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}
	total := 0
	n := len(s)

	for i := range n {
		current := hash[string(s[i])]
		if i < n-1 && current < hash[string(s[i+1])] {
			total -= current
			continue
		}

		total += current
	}

	return total
}
