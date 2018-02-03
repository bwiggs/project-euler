package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	in, _ := strconv.Atoi(os.Args[1])
	factor := solve(in)
	fmt.Printf("largest prime factor: %d\n", factor)
}

// solve is going to check all numbers k..n to see if they are divisible into n.
// to speed things up, we can reduce the number of k we need to check, by checking 2
// explicitely, then skipping all event numbers from k..n, thus divinding the
// amount of work by 2.
// we can further speed things up by
func solve(n int) int {

	lastFactor := 1
	f := 2

	if n%f == 0 {
		fmt.Printf("%d is factor\n", f)
		lastFactor = f
		n, f = reduce(n, f)
	}

	max := math.Sqrt(float64(n))

	f = 3
	for float64(f) <= max {
		if n%f == 0 {
			fmt.Printf("%d is a factor\n", f)
			lastFactor = f
			n, f = reduce(n, f)
		}
		f += 2
	}

	return lastFactor
}

func reduce(n, f int) (int, int) {
	n = n / f
	for n%f == 0 {
		n = n / f
	}
	return n, f
}
