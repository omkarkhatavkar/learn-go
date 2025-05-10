package main

import "testing"

// To run benchmarks, use the command:
// go test -bench=.
// The benchmark function name must start with `Benchmark` and take `*testing.B`.

// This benchmark will test the performance of our `Sum` function.
func BenchmarkSum(b *testing.B) {
	// Create a large slice to test with.
	input := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		input[i] = i
	}

	// b.N is automatically adjusted by the testing framework to get a stable benchmark result.
	// The loop runs b.N times, and the total time is measured.
	for i := 0; i < b.N; i++ {
		Sum(input)
	}

	// If the benchmark function does any setup that shouldn't be timed, call `b.ResetTimer()`.
	// b.ResetTimer()
	// // ... code to be timed
}
