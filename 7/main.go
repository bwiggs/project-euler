package main

import "fmt"

func main() {
	num := 1000
	p := brute(num)
	fmt.Printf("%d prime: %d\n", num, p)
}

func brute(num int) int {
	primes := []int{2}
	lastPrime := 2
	numPrimes := 1
	n := 3
	for {
		if numPrimes == num {
			break
		}
		// fmt.Printf("n = %d\n", n)
		if isPrime(n, primes) {
			// fmt.Printf("prime %d\n", n)
			primes = append(primes, n)
			lastPrime = n
			numPrimes++
		}
		n++
	}
	return lastPrime
}

func skipEvens(num int) int {
	primes := []int{2}
	lastPrime := 2
	numPrimes := 1
	n := 3
	for {
		if numPrimes == num {
			break
		}
		// fmt.Printf("n = %d\n", n)
		if isPrime(n, primes) {
			// fmt.Printf("prime %d\n", n)
			primes = append(primes, n)
			lastPrime = n
			numPrimes++
		}
		n += 2
	}
	return lastPrime
}

func isPrime(n int, primes []int) (isPrime bool) {
	isPrime = true
	for i := range primes {
		if n%primes[i] == 0 {
			isPrime = false
			break
		}
	}
	return
}
