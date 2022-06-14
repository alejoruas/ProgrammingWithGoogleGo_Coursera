package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p *Point) InitMe(xn, yn float64) {
	p.x = xn
	p.y = yn
}

func (p *Point) Scale(v float64) {
	p.x = p.x * v
	p.y = p.y * v
}

func (p *Point) Printme() {
	fmt.Println(p.x, p.y)
}

func (p Point) DistToOrig() float64 {
	t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
	return math.Sqrt(t)
}

func main() {
	p1 := Point{x: 3, y: 5}
	fmt.Println(p1.DistToOrig())

	var p Point
	p.InitMe(3, 4)
	p.Scale(4)
	p.Printme()

}
