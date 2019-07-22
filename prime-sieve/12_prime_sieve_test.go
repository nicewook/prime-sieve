package primesieve

import (
	"runtime"
	"testing"
)

func TestPrimeSieveConcurrency(t *testing.T) {
	PrimeSieveConcurrency()
}

func TestPrimeSieve(t *testing.T) {
	PrimeSieve()
}

func BenchmarkPrimeSieveConcurrency(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < b.N; i++ {
		PrimeSieveConcurrency()
	}
}

func BenchmarkPrimeSieve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeSieve()
	}
}
