package main

import (
	"testing"
)

// 1, 2, 3    [i-1, j-1] [i-1,j] [i+1, j+1]
// 4, 5, 6    _ [i,j] _
// 7, 8, 9   [i-1,j+1] [i+1, j] [i+2, j+1]

// func TestFibRec(t *testing.T) {
// 	res := fibRecursive(6)
// 	want := 8
// 	if res != want {
// 		t.Errorf("wrong result")
// 	}
// }

//	func TestFibIt(t *testing.T) {
//		res := fibIterative(6)
//		want := 8
//		if res != want {
//			t.Errorf("wrong result")
//		}
//	}
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

// //go test -bench=
