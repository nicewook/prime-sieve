package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func player(playerName string, table chan int) {
	for {
		// blocked till boll is coming.
		//then, ball hits the table anc comes to me
		ball := <-table

		ball++
		fmt.Printf("%s hit: %d\n", playerName, ball)
		time.Sleep(10 * time.Millisecond)
		table <- ball
	}
}

func main() {
	fmt.Printf("NumCPU: %d\n\n", runtime.NumCPU())
	// runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	var Ball int
	table := make(chan int)

	for i := 0; i < 100; i++ {
		go player(strconv.Itoa(i), table)
	}

	table <- Ball
	time.Sleep(1 * time.Second)
	<-table
}
