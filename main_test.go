package main

import (
	"testing"
)

func BenchmarkFibRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibRecursive(uint(40))
	}
}
func BenchmarkFibIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibIterative(uint(40))
	}
}
