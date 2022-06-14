package main

import "fmt"

type Shape2D interface {
	Area() float64
	Perimeter() float64
}

type Triangle struct{}

func (t Triangle) Area() float64 { return 0 }

func (t Triangle) Perimeter() float64 { return 0 }

type Rectangle struct{}

func (r Rectangle) Area() float64 { return 0 }

func (r Rectangle) Perimeter() float64 { return 0 }

func FitInYard(s Shape2D) bool {
	if s.Area() > 100 && s.Perimeter() > 100 {
		return true
	} else {
		return false
	}
}

func DrawShape(s Shape2D) bool {
	/*rect, ok := s.(Rectangle)
	if ok {
		DrawRect(rect)
	}
	tri, ok := s.(Triangle)
	if ok {
		DrawTriangle(tri)
	}*/

	switch sh := s.(type) {
	case Rectangle:
		DrawRect(sh)
	case Triangle:
		DrawTriangle(sh)
	default:
		return true
	}

	return false
}

func DrawRect(r Rectangle) bool    { return true }
func DrawTriangle(t Triangle) bool { return true }

/**********************************/

type Speaker interface{ Speak() }

type Dog struct{ name string }

/*func (d Dog) Speak() {
	fmt.Println(d.name)
}*/

func (d *Dog) Speak() {
	if d == nil {
		fmt.Println("<noise>")
	} else {
		fmt.Println(d.name)
	}
}

func main() {
	var s1 Speaker
	var d1 *Dog
	//d1.name = "Apolo"
	s1 = d1
	s1.Speak()
}
