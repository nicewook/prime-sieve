package primesieve

const maxPrimeNum = 104729 // the 10,000 th prime number is 104729

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

// PrimeSieveConcurrency uses golang concurrency
func PrimeSieveConcurrency() {
	ch := make(chan int)
	go generate(ch)

	for i := 0; i < 10000; i++ {
		prime := <-ch
		// fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

// PrimeSieve implemented without concurrency
func PrimeSieve() {
	num := make(map[int]bool, maxPrimeNum+1)
	numMax := int(maxPrimeNum ^ (1 / 2)) // max number for using as a sieve

	for i := 2; i < numMax+1; i++ {
		if num[i] == false {
			for j := i * i; j < maxPrimeNum; j += i {
				num[j] = true
			}
		}
	}

	// print
	for i := 2; i < maxPrimeNum; i++ {
		if num[i] == false {
			// fmt.Print(i, " ")
		}
	}
}
