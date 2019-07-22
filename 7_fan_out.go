package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(wID int, taskCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-taskCh
		if !ok { // channel closed
			return
		}
		d := time.Duration(task) * time.Millisecond
		time.Sleep(d)
		fmt.Printf("#p%d do task #%d\n", wID, task)
	}
}

func pool(wg *sync.WaitGroup, numWorkers, numTasks int) {
	taskCh := make(chan int)

	for i := 0; i < numWorkers; i++ {
		go worker(i, taskCh, wg)
	}

	for i := 0; i < numTasks; i++ {
		taskCh <- i
	}
	close(taskCh)
}

func main() {
	fmt.Println(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(36)
	go pool(&wg, 36, 50)
	wg.Wait()
}
