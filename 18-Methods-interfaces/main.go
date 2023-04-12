package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IntSlice []int

func (is IntSlice) String() string {
	var strs []string

	for _, v := range is {
		strs = append(strs, strconv.Itoa(v))
	}
	return "[" + strings.Join(strs, ";") + "]"
}

func main() {
	var v IntSlice = []int{1, 2, 3}
	var s fmt.Stringer = v // v is a stringer as String() method implemented and satisfied Stringer interface
	for i, x := range v {
		fmt.Printf("%d: %d\n", i, x)
	}

	fmt.Printf("%T %[1]v\n", v) // IntSlice basically implements Stringer Interface by adding String() method.
	fmt.Printf("%T %[1]v\n", s) // same behavior as expected.

}

//
//type Stringer interface {
//	String() string
//}
//
