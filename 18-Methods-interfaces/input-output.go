package main

import (
	"fmt"
	"io"
	"os"
)

//var i io.Writer

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	l := len(p)
	*b += ByteCounter(l)
	return l, nil
}

func main() {
	var c ByteCounter
	f1, _ := os.Open("in.txt")
	//f2, _ := os.Create("out.txt")
	f2 := &c

	n, _ := io.Copy(f2, f1) // Bytecounter implements write method, io.Copy() also expect a type satisfy write interface. where write metods returns bytes.

	fmt.Printf("Copied %d bytes\n", n)
	fmt.Println(c)
}
