package main

import "fmt"

func Add(x, y int) int {
	return x + y
}

func AddtoA(x int) func(int) int {
	return func(y int) int {
		return Add(x, y)
	}
}

func main() {
	add := AddtoA(3) // retuen "func(y int) int {
	//return Add(x, y)
	//}"
	// x is alreay there

	sum := add(2)
	fmt.Println(sum)
}
