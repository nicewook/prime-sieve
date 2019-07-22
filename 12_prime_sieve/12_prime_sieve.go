package primesieve

import (
	"fmt"
	"time"
)

const maxPrimeNum = 104729 // the 10,000th prime number is 104729

func generate(ch chan<- int) {
	for i := 2; i < maxPrimeNum+1; i++ {
		ch <- i
	}
}

func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

// Concurrency uses golang concurrency
func Concurrency() {
	nanos := time.Now().UnixNano()

	ch := make(chan int)
	go generate(ch)

	for i := 0; i < 10000; i++ {
		prime := <-ch
		if i < 10 {
			fmt.Print(prime, " ")
		}
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
	fmt.Println()
	fmt.Println(time.Now().UnixNano()-nanos, " ns")
}

// Sequential implemented without concurrency
func Sequential() {
	nanos := time.Now().UnixNano()
	num := make(map[int]bool, maxPrimeNum+1)
	numMax := int(maxPrimeNum ^ (1 / 2)) // max number for using as a sieve

	for i := 2; i < numMax+1; i++ {
		if num[i] == false {
			for j := i * i; j < maxPrimeNum+1; j += i {
				num[j] = true
			}
		}
	}

	// print the first 10
	cnt := 0
	for i := 2; i < maxPrimeNum+1; i++ {
		if num[i] == false && cnt < 10 {
			cnt++
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
	fmt.Println(time.Now().UnixNano()-nanos, " ns")
}
