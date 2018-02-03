package main

import "fmt"

func main() {
	fmt.Printf("sum: %d\n", solve(4000000))
}

func solve(max int) int {
	a, b, sum := 1, 2, 0
	for b <= max {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}
	return sum
}
