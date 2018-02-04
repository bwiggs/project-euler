package main

import "testing"

func benchSolve(num int, f func(int) ([]int, int), b *testing.B) {
	for n := 0; n < b.N; n++ {
		f(num)
	}
}

func BenchmarkN3100(b *testing.B)  { benchSolve(100, n3, b) }
func BenchmarkN31000(b *testing.B) { benchSolve(1000, n3, b) }

func BenchmarkN310000(b *testing.B) { benchSolve(10000, n3, b) }

func BenchmarkN2100(b *testing.B)    { benchSolve(100, n2, b) }
func BenchmarkN21000(b *testing.B)   { benchSolve(1000, n2, b) }
func BenchmarkN210000(b *testing.B)  { benchSolve(10000, n2, b) }
func BenchmarkN2100000(b *testing.B) { benchSolve(100000, n2, b) }
