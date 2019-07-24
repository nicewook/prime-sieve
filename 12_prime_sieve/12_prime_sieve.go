package primesieve

import (
	"fmt"
	"time"
)

// const maxPrimeNum = 104729 // the 10,000th prime number is 104729
// const primeCount = 10000

const maxPrimeNum = 1299827 // the 100,008th prime number is 1299827
const primeCount = 100008

func generate(ch chan<- int) {
	for i := 2; i < maxPrimeNum+1; i++ {
		ch <- i
	}
	close(ch)
}

func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i, ok := <-in
		if !ok {
			break
		}
		if i%prime != 0 {
			out <- i
			// fmt.Printf("%d ", i)
		}
	}
	// fmt.Println()
}

func filter2(in <-chan int, out chan<- int, prime int) {
	nextPrimeSquare := 0

	// search next prime
	for {
		i, ok := <-in
		if !ok {
			close(out)
			return
		}
		if i%prime != 0 {
			nextPrimeSquare = i * i
			out <- i
			break
		}
	}

	// sieve
	for {
		i, ok := <-in
		if !ok {
			close(out)
			return
		}

		if i < nextPrimeSquare {
			continue
		}
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

	for i := 0; i < primeCount; i++ {
		prime := <-ch
		if i < 10 {
			// fmt.Print(prime, " ")
		}
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
	fmt.Println()
	fmt.Println(time.Now().UnixNano()-nanos, " ns")
}

// Concurrency2 uses golang concurrency
func Concurrency2() {
	nanos := time.Now().UnixNano()

	ch := make(chan int)
	go generate(ch)

	for i := 0; i < primeCount; i++ {
		prime := <-ch
		ch1 := make(chan int)
		go filter2(ch, ch1, prime)
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
			// fmt.Print(i, " ")
		}
	}
	fmt.Println()
	fmt.Println(time.Now().UnixNano()-nanos, " ns")
}
