package main

import (
	"crypto/rand"
	"encoding/binary"
	"math"
	"sync"
)

// TrueRandom generates a cryptographically secure random float64 between 0 and 1
// Uses crypto/rand for true randomness instead of pseudo-random numbers
func TrueRandom() (float64, error) {
	// Read 8 random bytes from the cryptographically secure random number generator
	bytes := make([]byte, 8)
	_, err := rand.Read(bytes)
	if err != nil {
		return 0, err
	}

	// Convert the 8 bytes to a uint64 using big-endian byte order
	// This gives us a random integer between 0 and 2^64-1
	randomInt := binary.BigEndian.Uint64(bytes)

	// Normalize to [0, 1) by dividing by 2^64
	// We add 1 to MaxUint64 to ensure the result is never exactly 1.0
	return float64(randomInt) / float64(math.MaxUint64+1), nil
}

// TrueRandomRange generates a cryptographically secure random float64 between min and max
// Uses linear interpolation to scale the [0,1) range to [min, max)
func TrueRandomRange(min, max float64) (float64, error) {
	// Get a random value between 0 and 1
	r, err := TrueRandom()
	if err != nil {
		return 0, err
	}

	// Scale and shift the random value to the desired range
	// Formula: min + r * (max - min) maps [0,1) to [min, max)
	return min + r*(max-min), nil
}

// Sampling performs Monte Carlo sampling to test if a random point falls inside a unit circle
// This is the core of the Monte Carlo method for estimating π
// Returns true if the point (x,y) is inside the unit circle (x² + y² ≤ 1)
func Sampling() bool {
	// Generate random coordinates between 0 and 1
	// This represents a random point in the first quadrant of a unit square
	x, _ := TrueRandomRange(0, 1)
	y, _ := TrueRandomRange(0, 1)

	// Check if the point is inside the unit circle using the distance formula
	// A point is inside the circle if its distance from origin ≤ 1
	// We use x² + y² ≤ 1 instead of √(x² + y²) ≤ 1 to avoid expensive sqrt operation
	return x*x+y*y <= 1.0
}

// GenerateFromSampling runs n Monte Carlo samples concurrently to estimate π
// Uses goroutines for parallel sampling to improve performance
// Returns the count of samples that fell inside the unit circle
func GenerateFromSampling(n uint64) uint64 {
	// WaitGroup to synchronize all goroutines
	wg := sync.WaitGroup{}

	// Mutex to protect the shared counter from race conditions
	// Multiple goroutines will increment 'count' simultaneously
	mu := sync.Mutex{}

	// Counter for successful samples (points inside the circle)
	count := uint64(0)

	// Launch n goroutines, each performing one sample
	for range n {
		wg.Add(1) // Increment the WaitGroup counter

		go func() {
			defer wg.Done() // Decrement counter when goroutine completes

			// Perform one Monte Carlo sample
			if ok := Sampling(); ok {
				// Critical section: safely increment the shared counter
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}

	// Wait for all goroutines to complete
	wg.Wait()

	return count
}
