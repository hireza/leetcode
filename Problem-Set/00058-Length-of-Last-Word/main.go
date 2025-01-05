package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	splitS := strings.Split(s, " ")

	last := splitS[len(splitS)-1]

	return len(last)
}
