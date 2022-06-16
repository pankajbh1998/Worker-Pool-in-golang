package main

import (
	"fmt"
	"sync"
	"time"
)

// Consumer consumes tasks from jobs
func Consumer(trace int, n int, wg *sync.WaitGroup, jobs chan interface{}){
	fmt.Printf("Consumer %v started\n", trace)
	i := 0
	defer func(i *int){
		fmt.Printf("Consumer %v Finished %v task\n", trace, *i)
		wg.Done()
	}(&i)

	for job := range jobs {
		fmt.Printf("Consumer %v consumed :-> %v\n", trace, job)
		time.Sleep(time.Second*4)
		i++
	}
}
