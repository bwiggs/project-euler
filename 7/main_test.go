package main

import "testing"

func benchFindPrime(primeNum int, f func(int) int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		f(primeNum)
	}
}

func BenchmarkPrimeBruteForce10(b *testing.B)      { benchFindPrime(10, brute, b) }
func BenchmarkPrimeBruteForce100(b *testing.B)     { benchFindPrime(100, brute, b) }
func BenchmarkPrimeBruteForce1000(b *testing.B)    { benchFindPrime(1000, brute, b) }
func BenchmarkPrimeBruteForce10000(b *testing.B)   { benchFindPrime(10000, brute, b) }
func BenchmarkPrimeBruteForce100000(b *testing.B)  { benchFindPrime(100000, brute, b) }
func BenchmarkPrimeBruteForce1000000(b *testing.B) { benchFindPrime(1000000, brute, b) }

func BenchmarkPrimeSkipEvens10(b *testing.B)      { benchFindPrime(10, skipEvens, b) }
func BenchmarkPrimeSkipEvens100(b *testing.B)     { benchFindPrime(100, skipEvens, b) }
func BenchmarkPrimeSkipEvens1000(b *testing.B)    { benchFindPrime(1000, skipEvens, b) }
func BenchmarkPrimeSkipEvens10000(b *testing.B)   { benchFindPrime(10000, skipEvens, b) }
func BenchmarkPrimeSkipEvens100000(b *testing.B)  { benchFindPrime(100000, skipEvens, b) }
func BenchmarkPrimeSkipEvens1000000(b *testing.B) { benchFindPrime(1000000, skipEvens, b) }
