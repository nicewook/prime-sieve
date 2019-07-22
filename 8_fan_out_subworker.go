package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	WORKERS    = 5
	SUBWORKERS = 3
	TASKS      = 20
	SUBTASKS   = 10
)

func subworker(tID, sID int, sTask chan int) {
	for {
		task, ok := <-sTask
		if !ok { // subtasks channel closed
			return
		}
		time.Sleep(time.Duration(task) * time.Millisecond)
		fmt.Printf("P%d, S%d: task: %d\n", tID, sID, task)
	}
}

func worker(wID int, taskCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-taskCh
		if !ok { // channel closed
			return
		}

		subtasks := make(chan int)
		for i := 0; i < SUBWORKERS; i++ {
			go subworker(wID, i, subtasks)
		}
		for i := 0; i < SUBTASKS; i++ {
			subtasks <- (task * i)
		}

		// time.Sleep(3 * time.Second)
		close(subtasks)
	}
}

func main() {
	fmt.Println(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(WORKERS)
	tasks := make(chan int)

	for i := 0; i < WORKERS; i++ {
		go worker(i, tasks, &wg)
	}

	for i := 0; i < TASKS; i++ {
		tasks <- i
	}
	close(tasks)
	wg.Wait()
}
