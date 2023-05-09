package main

import (
	"log"
	"time"
)

func main() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	for i := range chans {
		go func(i int, ch chan<- int) {
			for {
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i
			}
		}(i+1, chans[i])
	}

	for i := 0; i < 12; i++ {
		select {
		case m1 := <-chans[0]:
			log.Println("Received", m1)

		case m2 := <-chans[1]:
			log.Println("Received", m2)
		}
	}
}
