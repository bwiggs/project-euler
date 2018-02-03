package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("reading input: %v", err)
		os.Exit(1)
	}
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
		if isPrime(n) {
			// if isPrimeWithKnownPrimes(n, primes) {
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
		if isPrime(n) {
			// if isPrimeWithKnownPrimes(n, primes) {
			primes = append(primes, n)
			lastPrime = n
			numPrimes++
		}
		n += 2
	}
	return lastPrime
}

func isPrimeWithKnownPrimes(n int, primes []int) (isPrime bool) {
	isPrime = true
	max := math.Sqrt(float64(n)) // short cut for checking primes
	for i := range primes {
		if float64(primes[i]) >= max {
			break
		}
		if n%primes[i] == 0 {
			isPrime = false
			break
		}
	}
	return
}

func isPrime(n int) bool {

	if n == 1 {
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
		f += 6
	}
	return true
}
