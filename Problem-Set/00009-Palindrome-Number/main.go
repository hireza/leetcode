package main

func isPalindrome(x int) bool {
	num := 0

	if x < 0 {
		return false
	}

	temp := x
	for temp > 0 {
		num = num*10 + temp%10
		temp /= 10
	}

	return num == x
}
