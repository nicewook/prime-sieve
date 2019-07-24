package primesieve

import (
	"testing"
)

func TestConcurrency(t *testing.T) {
	Concurrency()
}

func TestConcurrency2(t *testing.T) {
	Concurrency2()
}

func TestSequential(t *testing.T) {
	Sequential()
}

// func BenchmarkPrimeSieveConcurrency(b *testing.B) {
// 	runtime.GOMAXPROCS(runtime.NumCPU())
// 	for i := 0; i < b.N; i++ {
// 		Concurrency()
// 	}
// }

// func BenchmarkPrimeSieve(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Sequential()
// 	}
// }
