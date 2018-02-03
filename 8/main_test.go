package main

import "testing"

func benchSolveProducts(numProducts int, f func() string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		f()
	}
}

func BenchmarkStringBased4(b *testing.B)  { benchSolveProducts(4, strBased, b) }
func BenchmarkStringBased6(b *testing.B)  { benchSolveProducts(6, strBased, b) }
func BenchmarkStringBased8(b *testing.B)  { benchSolveProducts(8, strBased, b) }
func BenchmarkStringBased10(b *testing.B) { benchSolveProducts(10, strBased, b) }
func BenchmarkStringBased13(b *testing.B) { benchSolveProducts(13, strBased, b) }
func BenchmarkStringBased20(b *testing.B) { benchSolveProducts(20, strBased, b) }

func BenchmarkDigitBased4(b *testing.B)  { benchSolveProducts(4, digitBased, b) }
func BenchmarkDigitBased6(b *testing.B)  { benchSolveProducts(6, digitBased, b) }
func BenchmarkDigitBased8(b *testing.B)  { benchSolveProducts(8, digitBased, b) }
func BenchmarkDigitBased10(b *testing.B) { benchSolveProducts(10, digitBased, b) }
func BenchmarkDigitBased13(b *testing.B) { benchSolveProducts(13, digitBased, b) }
func BenchmarkDigitBased20(b *testing.B) { benchSolveProducts(20, digitBased, b) }
