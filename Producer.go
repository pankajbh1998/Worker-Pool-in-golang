package main

import (
	"fmt"
	"sync"
	"time"
)
// Producer produces n tasks or max tasks till time tick
func Producer(trace int, n int,wg *sync.WaitGroup,  jobs chan interface{}){
	fmt.Printf("Producer %v started\n", trace)
	var i int
	defer func(i *int){
		fmt.Printf("Producer %v Finished %v Task \n", trace, *i)
		wg.Done()
	}(&i)

	tick := time.Tick(time.Second*40)
	for i<n {
		select {
		case <- tick:
			return
		default:
			rand := GetRandTask()
			jobs <- rand
			fmt.Printf("Producer %v pushed :-> %v\n", trace, rand)
			time.Sleep(time.Second)
			i++
		}
	}
}