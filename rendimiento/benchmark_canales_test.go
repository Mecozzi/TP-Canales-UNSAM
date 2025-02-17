package main

import (
	"testing"
)

// Benchmark: Canal con búfer
func BenchmarkCanalConBuffer(b *testing.B) {
	canal := make(chan int, 100)

	go func() {
		for i := 0; i < b.N; i++ {
			canal <- i
		}
		close(canal)
	}()

	for range canal {
	}
}
