package main

import (
	"fmt"
	"sync"
)

func do() int {
	var value int
	var w sync.WaitGroup

	m := make(chan bool, 1)

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go func() {
			m <- true
			value++ // race condition
			<-m
			w.Done()
		}()
	}
	w.Wait()
	return value
}

func main() {
	fmt.Println(do())
}
