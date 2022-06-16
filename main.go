package main

import (
	"fmt"
	"sync"
)

func main(){
	// total number of workers required
	numberOfProducer := 5
	numberOfConsumer := 3
	fmt.Printf("Queue started with %v producers and %v consumers\n",numberOfProducer, numberOfConsumer)

	//wait for the workers to finish the tasks
	wgProducer := &sync.WaitGroup{}
	wgConsumer := &sync.WaitGroup{}
	// jobs produced by producer worker remain in jobs channel until not consumed by consumer worker
	jobs := make(chan interface{}, 20)

	// start the worker
	startWorker(numberOfProducer, wgProducer, jobs, Producer)
	startWorker(numberOfConsumer, wgConsumer, jobs, Consumer)

	// wait for producers to publish the tasks
	wgProducer.Wait()
	close(jobs)
	// wait for consumers to finish the entire task
	wgConsumer.Wait()

	fmt.Println("Tasks completed successfully")
}

// startWorker will start the n workers
func startWorker(n int, wg *sync.WaitGroup, jobs chan interface{}, f func(int, int,*sync.WaitGroup,chan interface{})){
	for i:=1;i<=n;i++ {
		wg.Add(1)
		go f(i, n+i, wg, jobs)
	}
}
