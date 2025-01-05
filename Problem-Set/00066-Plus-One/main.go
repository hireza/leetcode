package main

func plusOne(digits []int) []int {
	temp := 0
	for i := len(digits) - 1; i >= 0; i-- {
		total := digits[i]
		if i == len(digits)-1 {
			total++
		}

		if temp == 1 {
			total += temp
			temp = 0
		}

		if total == 10 {
			temp = 1
		}

		digits[i] = total % 10
	}

	if temp == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
