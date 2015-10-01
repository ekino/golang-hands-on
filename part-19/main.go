package main

import (
	"fmt"
	"time"
	"sync"
)

type Job struct {
	Command string
}

var wg sync.WaitGroup

func worker(id int, queue chan *Job) {
	for job := range queue {
		fmt.Printf("#%d - Cmd: %s\n", id, job.Command)
		time.Sleep(5000 * time.Millisecond)
		wg.Done()
	}
}

func main() {
	queue := make(chan *Job, 5)

	for i := 0; i < 5; i++ {
		go worker(i, queue)
	}

	for i := 0; i < 50; i++ {
		wg.Add(1)

		fmt.Printf("Send cmd: run %d\n", i)
		job := &Job{fmt.Sprintf("run %d", i)}

		queue <- job
	}

	wg.Wait()

	fmt.Println("done")
}