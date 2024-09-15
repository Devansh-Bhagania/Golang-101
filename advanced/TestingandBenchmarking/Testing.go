// Testing ensures your code works as expected, and benchmarking helps identify performance bottlenecks.

package math

import "testing"

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

// Benchmarking

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

// Profiling
// Use the pprof tool to analyze CPU and memory usage.

// $ go test -cpuprofile cpu.out -memprofile mem.out
// $ go tool pprof cpu.out
