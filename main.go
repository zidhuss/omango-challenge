package main

import (
	"fmt"
	"math"
	"strconv"
)

// Solve returns true if the number is a K number
func Solve(number int) bool {
	// Write your code here

	if number == 1 {
		return true
	}

	numberSquared := fmt.Sprintf("%d", int(math.Pow(float64(number), 2)))

	for i := 0; i < len(numberSquared)-1; i++ {

		num1, _ := strconv.Atoi(numberSquared[0 : i+1])
		num2, _ := strconv.Atoi(numberSquared[i+1 : len(numberSquared)])

		if num1 == 0 || num2 == 0 {
			continue
		}

		if num1+num2 == number {
			return true
		}

	}

	return false
}

// 0.685s
