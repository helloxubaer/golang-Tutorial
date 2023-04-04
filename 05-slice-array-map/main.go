package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Println(len(words), "Unique words")
	fmt.Println(words)

	type kv struct {
		key   string
		value int
	}

	var ss []kv

	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].value > ss[j].value
	})

	fmt.Println("ss: ", ss)

	for _, s := range ss {
		fmt.Println(s.key, "appears", s.value, "times")
	}

}
