package main

import "fmt"

func main() {
	num := 1000
	trips, prod := n3(num)
	fmt.Printf("trips: %v, product: %d", trips, prod)
}

// my initial attempt. The third loop was dumb. I already had the only value C
// could be, so I could have just removed it and used the 2 loops for a cleaner
// brute force solution.
func n3(num int) ([]int, int) {
	for a := 1; a < num-2; a++ {
		for b := a + 1; b < num-1; b++ {
			for c := num - b - a; c > b; c-- {
				if a+b+c == num && a*a+b*b == c*c {
					// fmt.Printf("%3d + %3d + %3d = %d\n", a, b, c, num)
					// fmt.Printf("(%3d^2, %3d^2 =  %3d^2)\n", a, b, c)
					// fmt.Printf("%3d * %3d * %3d = %d\n", a, b, c, a*b*c)
					return []int{a, b, c}, a * b * c
				}
			}
		}
	}
	return nil, 0
}
