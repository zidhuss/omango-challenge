package main

import "strconv"

// Solve returns true if the number is a K number
func Solve(number int) bool {
	// Write your code here

	sqr := number * number

	sqrString := strconv.Itoa(sqr)

	for i := 0; i < len(sqrString); i++ {

		left, _ := strconv.Atoi(sqrString[0:i])
		right, _ := strconv.Atoi(sqrString[i:])

		if left+right == number && right != 0 {
			return true
		}
	}
	return false

}
