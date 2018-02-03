package main

import "testing"

func benchFindPrime(primeNum int, f func(int) int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		f(primeNum)
	}
}

func BenchmarkPrimeBruteForce100(b *testing.B)   { benchFindPrime(100, brute, b) }
func BenchmarkPrimeBruteForce1000(b *testing.B)  { benchFindPrime(1000, brute, b) }
func BenchmarkPrimeBruteForce10000(b *testing.B) { benchFindPrime(10000, brute, b) }
func BenchmarkPrimeBruteForce20000(b *testing.B) { benchFindPrime(20000, brute, b) }
func BenchmarkPrimeBruteForce30000(b *testing.B) { benchFindPrime(30000, brute, b) }

func BenchmarkPrimeSkipEvens100(b *testing.B)   { benchFindPrime(100, skipEvens, b) }
func BenchmarkPrimeSkipEvens1000(b *testing.B)  { benchFindPrime(1000, skipEvens, b) }
func BenchmarkPrimeSkipEvens10000(b *testing.B) { benchFindPrime(10000, skipEvens, b) }
func BenchmarkPrimeSkipEvens20000(b *testing.B) { benchFindPrime(20000, skipEvens, b) }
func BenchmarkPrimeSkipEvens30000(b *testing.B) { benchFindPrime(30000, skipEvens, b) }
