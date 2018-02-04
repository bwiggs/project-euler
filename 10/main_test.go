package main

import "testing"

func benchSolve(num int, f func(int) int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		f(num)
	}
}

func BenchmarkLinear1000(b *testing.B)    { benchSolve(1000, linear, b) }
func BenchmarkLinear10000(b *testing.B)   { benchSolve(10000, linear, b) }
func BenchmarkLinear100000(b *testing.B)  { benchSolve(100000, linear, b) }
func BenchmarkLinear1000000(b *testing.B) { benchSolve(1000000, linear, b) }

func BenchmarkLinearSmart1000(b *testing.B)    { benchSolve(1000, linearSmart, b) }
func BenchmarkLinearSmart10000(b *testing.B)   { benchSolve(10000, linearSmart, b) }
func BenchmarkLinearSmart100000(b *testing.B)  { benchSolve(100000, linearSmart, b) }
func BenchmarkLinearSmart1000000(b *testing.B) { benchSolve(1000000, linearSmart, b) }
