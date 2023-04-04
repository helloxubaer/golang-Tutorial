package main

import "fmt"

func fib() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return b
	}
}

func main() {
	f := fib() // as long as this f exists, a & b exist.

	for x := f(); x < 100; x = f() {
		fmt.Println(x)
	}

	// example:
	fmt.Println("=======================")
	s := make([]func(), 4)

	for i := 0; i < 4; i++ {
		i2 := i // closure capture
		s[i] = func() {
			fmt.Printf("%d @ %p\n", i2, &i2)
		}
	}

	for i := 0; i < 4; i++ {
		s[i]()
	}

}
