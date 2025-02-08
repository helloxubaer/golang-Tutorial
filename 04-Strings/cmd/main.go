package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := "Über"

	fmt.Printf("%8T %[1]v\n", s)
	fmt.Printf("%8T %[1]v\n", []rune(s))
	b := []byte(s)
	fmt.Printf("%8T %[1]v %d\n", b, len(b))

	fmt.Println("---------------------")
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Not enough args")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		fmt.Println(s)
		t := strings.Join(s, new)
		fmt.Println(t)
	}

}
