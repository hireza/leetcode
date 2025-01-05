package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	maxLen := len(a)
	if len(a) < len(b) {
		maxLen = len(b)
	}

	a = fmt.Sprintf("%0*s", maxLen, a)
	b = fmt.Sprintf("%0*s", maxLen, b)

	temp := 0
	res := ""
	for i := maxLen - 1; i >= 0; i-- {
		iterA, _ := strconv.Atoi(string(a[i]))
		iterB, _ := strconv.Atoi(string(b[i]))

		tot := iterA + iterB + temp
		if tot >= 2 {
			temp = 1
		} else if tot == 1 {
			temp = 0
		} else {
			temp = 0
		}

		res = strconv.Itoa(tot%2) + res
	}

	if temp == 1 {
		res = "1" + res
	}

	return res
}
