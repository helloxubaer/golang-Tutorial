package main

import (
	"fmt"
	"net/http"
)

//var nextID int

// var nextID = make(chan int) // even better way is a type nextID and bind method hanldler on that type

type nextID chan int

func (ch nextID) handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>You got %d<h1>", <-ch)
	// nextID++ // very unsafe
}

func counter(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func main() {
	var nextID nextID = make(chan int)
	go counter(nextID)
	http.HandleFunc("/", nextID.handler)
	http.ListenAndServe(":8080", nil)
}
