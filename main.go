package main

import (
	"fmt"
	"strconv"
)

func main() {
	//printing all K number under 1,000,000
	for i := 0; i < 1000000; i++ {
		if Solve(i) {
			fmt.Println(i)
		}
	}
}

//better performance only max of 3 checks for a number :)
func Solve(n int) bool {
	l := len(strconv.Itoa(n))
	s := strconv.Itoa(n * n)
	for i := max(l-2, 0); i <= l; i++ {
		a, _ := strconv.Atoi(s[0:i])
		b, _ := strconv.Atoi(s[i:])
		if b != 0 && a+b == n {
			return true
		}
	}
	return false
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
