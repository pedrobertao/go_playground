// Package main demonstrates Go iterator patterns using the iter package.
// This program showcases both built-in iterators from maps/slices packages
// and custom iterator implementation.
package main

import (
	"fmt"    // for formatted output
	"iter"   // for iterator types (Go 1.23+)
	"maps"   // for map iterator utilities
	"slices" // for slice iterator utilities
)

// Iterator type definitions for reference:
// Seq[V any] represents an iterator over sequences of individual values
// Seq2[K, V any] represents an iterator over sequences of key-value pairs

// CountMax creates a custom iterator that yields integers from 0 to max (inclusive).
// This demonstrates how to implement the iter.Seq[int] interface.
//
// Parameters:
//
//	max: the maximum value to count up to (inclusive)
//
// Returns:
//
//	An iterator that can be used with range loops
//
// Example usage:
//
//	for i := range CountMax(5) {
//	    fmt.Println(i) // prints 0, 1, 2, 3, 4, 5
//	}
func CountMax(max int) iter.Seq[int] {
	// Return a function that implements the iterator protocol
	return func(yield func(int) bool) {
		// Iterate from 0 to max (inclusive)
		for i := 0; i <= max; i++ {
			// Call yield with the current value
			// If yield returns false, stop iteration (early termination)
			if !yield(i) {
				return
			}
		}
	}
}

// main demonstrates various iterator patterns in Go 1.23+
func IterExampl() {
	// Create empty collections for demonstration
	myMap := make(map[string]int) // empty map
	mySlice := make([]int, 0)     // empty slice

	// Example 1: Iterate over map using maps.All()
	// maps.All returns an iterator that yields key-value pairs
	// Note: No output since myMap is empty
	for k, v := range maps.All(myMap) {
		fmt.Printf("%s: %v\n", k, v)
	}

	// Example 2: Iterate over slice using slices.All()
	// slices.All returns an iterator that yields index-value pairs
	// Note: No output since mySlice is empty
	for i, val := range slices.All(mySlice) {
		fmt.Printf("[%d]: %v\n", i, val)
	}

	// Example 3: Iterate over slice values only using slices.Values()
	// slices.Values returns an iterator that yields only values (no indices)
	// Note: No output since mySlice is empty
	for val := range slices.Values(mySlice) {
		fmt.Println(val)
	}

	// Example 4: Use our custom CountMax iterator
	// This will print numbers from 0 to 10 (inclusive)
	for digit := range CountMax(10) {
		fmt.Println(digit) // prints: 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
	}
}
