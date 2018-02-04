package main

import (
	"fmt"
	"math"
)

func main() {
	num := 1000
	sum := linear(num)
	fmt.Printf("Prime Sum: %d\n", sum)

	sum = sieve(num)
	fmt.Printf("Prime Sum: %d\n", sum)
}

func sieve(n int) int {
	nums := make([]bool, n-1)

	// mark the 2s
	// idx 2 is the first event number, skipping 2
	for i := 2; i < n-1; i += 2 {
		nums[i] = true
	}

	// mak the 3s
	// idx 5 is the first multiple of 3
	for i := 4; i < n-1; i += 3 {
		nums[i] = true
	}

	for i := 3; i < n-1; i++ {
		if !nums[i] {
			inc := i + 2
			for j := i + inc; j < n; j += inc {
				// fmt.Printf("i: %3d j: %3d inc: %3d\n", i, j, inc)
				nums[i] = true
			}
		}
	}

	sum := 0
	for i := range nums {
		fmt.Printf("v: %5d v:%v\n", i+2, nums[i])
		if !nums[i] {
			sum += i + 2
		}
	}
	return sum
}

func linearSmart(n int) int {
	i, sum := 5, 5
	for {
		if i >= n {
			return sum
		}

		if isPrime(i) {
			sum += i
		}

		i += 2

		if isPrime(i) {
			sum += i
		}

		i += 4
	}
}

func linear(n int) int {
	sum := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			// fmt.Printf("%d ", i)
			sum += i
		}
	}
	return sum
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n < 4 {
		return true
	} else if n%2 == 0 {
		return false
	} else if n < 9 {
		return true
	} else if n%3 == 0 {
		return false
	}

	r := int(math.Floor(math.Sqrt(float64(n))))
	f := 5
	for f <= r {
		if n%f == 0 {
			return false
		} else if n%(f+2) == 0 {
			return false
		}
		// n%(f+2) didnt actually advance the pointer, so we need to do that here.
		f += 6
	}
	return true

}
