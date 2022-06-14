package main

import "fmt"

func main() {
	var a float64
	var v float64
	var so float64
	var t float64

	fmt.Println("Please enter the acceleration")
	fmt.Scan(&a)
	fmt.Println("Please enter the initial velocity")
	fmt.Scan(&v)
	fmt.Println("Please enter the initial displacement")
	fmt.Scan(&so)

	fn := GenDisplaceFn(a, v, so)

	fmt.Println("Please enter the time")
	fmt.Scan(&t)

	s := fn(t)

	fmt.Printf("The displacement is %f\n", s)
}

func GenDisplaceFn(a float64, v float64, so float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*(t*t) + v*t + so
	}
	return fn
}
