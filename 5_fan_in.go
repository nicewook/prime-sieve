package main

import (
	"fmt"
	"time"
)

func producer(pName string, ch chan string, d time.Duration) {
	var i int
	for {
		ch <- fmt.Sprintf("%s: %d", pName, i)
		i++
		time.Sleep(d)
	}
}

func reader(out chan string) {
	for x := range out {
		fmt.Println(x)
	}
}

func main() {
	ch := make(chan string)
	out := make(chan string)

	go producer("pd1", ch, 100*time.Millisecond)
	go producer("pd2", ch, 250*time.Millisecond)
	go reader(out)

	for i := range ch {
		out <- i
	}
}
