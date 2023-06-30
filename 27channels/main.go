package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Learning about channels")

	myCh := make(chan int,2)

	wg := &sync.WaitGroup{}

	// this is an error because we cannot access channels outside of goroutines
	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)
	// receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val,isChannelOpen := <-myCh
		// the above code is used to avoid a race condition
		// if no value is written to a channel, it will show 0
		// but what if we send the signal 0
		// to avoid that case we use the above code
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	// send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		// myCh <- 6
		// close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
