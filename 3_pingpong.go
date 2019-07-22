package main

import (
	"fmt"
	"time"
)

func player(playerName string, table chan int) {
	for {
		// blocked till boll is coming.
		//then, ball hits the table anc comes to me
		ball := <-table

		ball++
		fmt.Printf("%s hits the ball: %d\n", playerName, ball)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

func main() {
	var Ball int
	table := make(chan int)
	go player("Alice", table)
	go player("Bob", table)

	table <- Ball
	time.Sleep(1 * time.Second)
	<-table
}
