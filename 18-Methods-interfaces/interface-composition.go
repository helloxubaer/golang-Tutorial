package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

type Line struct {
	Begin, End Point
}

type Distancer interface {
	Distance() float64
}

func PrintDistancer(d Distancer) {
	fmt.Println(d.Distance())
}

type Path []Point

func (l Line) Distance() float64 {
	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
}

func (l *Line) ScalBy(f float64) {
	l.End.X += (f - 1) * (l.End.X - l.Begin.X)
	l.End.Y += (f - 1) * (l.End.Y - l.Begin.Y)
}

func (p Path) Distance() float64 {
	var sum float64
	for i := 1; i < len(p); i++ {
		sum += Line{p[i-1], p[i]}.Distance()
	}
	return sum
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {

	side := Line{Point{1, 2}, Point{4, 6}}
	perimeter := Path{{1, 2}, {2, 3}}
	Distancer := side
	fmt.Println(side.Distance())
	fmt.Println(perimeter.Distance())
	fmt.Println(Distancer.Distance()) // also works
	fmt.Println(Distancer.Distance()) // also works

	PrintDistancer(side) // also works

	//fmt.Println(Line{Point{1, 2}, Point{4, 6}}.ScalBy(2.0)) // will not work as receiver is a pointer. without pointer you can return the modified value and use it.

	p, q := Point{1, 1}, ColoredPoint{Point{5, 4}, color.RGBA{255, 0, 0, 255}}

	l1 := q.Distance(p)
	l2 := p.Distance(q) // no as q is a colored point. but first modifiy distance method to take point argument

	fmt.Println(l1, l2)

}
