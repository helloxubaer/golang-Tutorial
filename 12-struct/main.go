package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

func main() {

	c := map[string]*Employee{}

	c["Zubaer"] = &Employee{"Zubaer", 1, nil, time.Now()}
	//c["Zubaer"].Number++ # will not work if address of an entry can change in map.

	c["Maria"] = &Employee{"Maria", 2, c["Zubaer"], time.Now()}

	fmt.Printf("%T, %[1]v\n", c["Zubaer"])
	fmt.Printf("%T, %[1]v\n", c["Maria"])

}
