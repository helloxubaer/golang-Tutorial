package main

import (
	"fmt"
	"os"
)

func main() {
	var sum float64
	var n int

	for {
		var value float64
		_, err := fmt.Fscanln(os.Stdin, &value)
		if err != nil {
			break
		}
		sum += value
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}
	fmt.Printf("Average is: %v", sum)
}
