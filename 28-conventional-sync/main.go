package main

import (
	"fmt"
	"sync"
)

func do() int {
	var value int
	var w sync.WaitGroup

	// m := make(chan bool, 1) // counting semaphore, works as mutex
	// better to use sync.Mutex
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go func() {
			// m <- true // write and block until read
			m.Lock()
			value++ // race condition
			// <-m       // read
			m.Unlock()
			w.Done()
		}()
	}
	w.Wait()
	return value
}

func main() {
	fmt.Println(do())
}
