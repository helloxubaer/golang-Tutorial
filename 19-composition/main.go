package main

import (
	"fmt"
	"sort"
)

type Organ struct {
	Name   string
	Weight int
}

type Organs []Organ

func (o Organs) Len() int {
	return len(o)
}

func (o Organs) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

type ByName struct{ Organs }
type ByWeight struct{ Organs }

func (s ByName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}

func (s ByWeight) Less(i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

func main() {
	organs := []Organ{{"brain", 1340}, {"liver", 1494}, {"apleen", 162}, {"heart", 290}}
	var ByName = ByName{organs}
	fmt.Println(organs)

	sort.Sort(ByName)
	fmt.Println(organs)
	sort.Sort(ByWeight{organs})
	fmt.Println(organs)

}
