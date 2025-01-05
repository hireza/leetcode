package main

func climbStairs(n int) int {
	first, second := 0, 1
	for n > 0 {
		first, second = second, first+second
		n--
	}
	return second
}
