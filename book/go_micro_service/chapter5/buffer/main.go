package main

import (
	"fmt"
	"sync"
)

const (
	noGoroutine = 5
	noTask      = 10
)

var wg sync.WaitGroup

func main() {
	tasks := make(chan int, noTask)
	for no := 1; no <= noGoroutine; no++ {
		wg.Add(1)
		go taskProcess(tasks, no)
	}

	//向tasks缓存通道内放入任务号
	for taskNO := 1; taskNO <= noTask; taskNO++ {
		tasks <- taskNO
	}
	close(tasks)
	wg.Wait()
}

func taskProcess(tasks chan int, workerNo int) {
	defer wg.Done()
	for t := range tasks {
		fmt.Printf("Worker %d is processing Task no:%d\n", workerNo, t)
	}
	fmt.Printf("Worker %d got off work \n", workerNo)
}
