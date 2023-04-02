package main

import (
	"fmt"
	"hello"
	"os"
)

func main() {

	if len(os.Args) > 1 {
		fmt.Printf(hello.SayHello(os.Args[1]))
	} else {
		fmt.Println("Hello world from golang!")
	}

}
