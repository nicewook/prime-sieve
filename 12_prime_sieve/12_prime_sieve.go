package primesieve

import (
	"fmt"
	"time"
)

const maxPrimeNum = 104729 // the 10,000th prime number is 104729
const primeCount = 10000

// const maxPrimeNum = 1299827 // the 100,008th prime number is 1299827
// const primeCount = 100008

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

func filter2(in <-chan int, out chan<- int, prime int) {
	primeSquare := prime * prime

	// skip smaller than the square of the prime
	for {
		if i := <-in; i < primeSquare {
			out <- i
		} else {
			break
		}
	}

	// sieve
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

// Concurrency uses golang concurrency
func Concurrency() []int {
	nanos := time.Now().UnixNano()
	var result []int

	ch := make(chan int)
	go generate(ch)

	for i := 0; i < primeCount; i++ {
		prime := <-ch
		result = append(result, prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
	fmt.Println("Concurrency elapsed time: ", time.Now().UnixNano()-nanos, " ns")
	return result
}

// Concurrency2 uses golang concurrency
func Concurrency2() []int {
	nanos := time.Now().UnixNano()
	var result []int

	ch := make(chan int)
	go generate(ch)

	for i := 0; i < primeCount; i++ {
		prime := <-ch
		result = append(result, prime)
		ch1 := make(chan int)
		go filter2(ch, ch1, prime)
		ch = ch1
	}
	fmt.Println("Concurrency2 elapsed time: ", time.Now().UnixNano()-nanos, " ns")
	return result
}

// Sequential implemented without concurrency
func Sequential() []int {
	nanos := time.Now().UnixNano()
	var result []int

	num := make(map[int]bool, maxPrimeNum+1)
	numMax := int(maxPrimeNum ^ (1 / 2)) // max number for using as a sieve

	for i := 2; i < numMax+1; i++ {
		if num[i] == false {
			for j := i * i; j < maxPrimeNum+1; j += i {
				num[j] = true
			}
		}
	}

	// print
	for i := 2; i < maxPrimeNum+1; i++ {
		if num[i] == false {
			result = append(result, i)
		}
	}
	fmt.Println("Sequential elapsed time: ", time.Now().UnixNano()-nanos, " ns")
	return result
}
