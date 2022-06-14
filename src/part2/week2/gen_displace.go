package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
}

func main() {
	var sa, sv0, ss0, st string
	fmt.Println("Please enter the value of acceleration (a): ")
	fmt.Scan(&sa)
	a, err := strconv.ParseFloat(sa, 64)
	if err != nil {
		log.Println("Get input number failed:\n", err)
		return
	}
	fmt.Println("Please enter the value of initial velocity (v0): ")
	fmt.Scan(&sv0)
	v0, err := strconv.ParseFloat(sv0, 64)
	if err != nil {
		log.Println("Get input number failed:\n", err)
		return
	}
	fmt.Println("Please enter the value of initial displacement (s0): ")
	fmt.Scan(&ss0)
	s0, err := strconv.ParseFloat(ss0, 64)
	if err != nil {
		log.Println("Get input number failed:\n", err)
		return
	}
	fn := GenDisplaceFn(a, v0, s0)
	for {
		fmt.Println("Please enter the value of time (t), or 'X' to exit: ")
		fmt.Scan(&st)
		if st == "X" {
			break
		}
		t, err := strconv.ParseFloat(st, 64)
		if err != nil {
			log.Println("Get input number failed:\n", err)
			return
		}
		if t < 0 {
			fmt.Println("Invalid time: should be non-negative!")
			continue
		}
		fmt.Println("The displacement is ", fn(t))
	}
}
