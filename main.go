package main

import (
	"strconv"
)

// Solve returns true if the number is a K number
func Solve(number int) bool {
	n := number * number
	if n < 10 {
		return true
	}

	s := strconv.Itoa(n)
	for i := 1; i < len(s); i++ {
		x, _ := strconv.Atoi(s[0:i])
		y, _ := strconv.Atoi(s[i:])

		if x > number {
			return false
		}

		if y == 0 || y > number {
			continue
		}

		if x+y == number {
			return true
		}
	}

	return false
}
