package primesieve

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func TestConcurrency(t *testing.T) {
	r1 := Concurrency()
	r2 := Concurrency2()
	r3 := Sequential()
	fmt.Printf("length r1, r2, r3: %d, %d, %d\n", len(r1), len(r2), len(r3))
	if !reflect.DeepEqual(r1, r2) {
		t.Error("Compare r1, r2: Not same!")
	}

	if !reflect.DeepEqual(r2, r3) {
		t.Error("Compare r2, r3: Not same!")
	}

	r1[0] = 0xFF
	if reflect.DeepEqual(r1, r2) {
		t.Error("Compare modified r1, r2: Shouldn't be the same!")
	}

}

func BenchmarkPrimeSieveConcurrency(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < b.N; i++ {
		Concurrency()
	}
}

func BenchmarkPrimeSieveConcurrency2(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < b.N; i++ {
		Concurrency2()
	}
}

func BenchmarkPrimeSieve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sequential()
	}
}
