package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world from golang!")
	fmt.Printf("Hello, %s\n", os.Args[1])
}
