package fib

import "testing"

// from fib_test.go
func BenchmarkFib35(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fib(35)
	}
}

func BenchmarkFibMemo35(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		memo := make([]int, 35+1)
		fibMemo(35, memo)
	}
}

func BenchmarkFibBottomUp35(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibBottomUp(35)
	}
}
