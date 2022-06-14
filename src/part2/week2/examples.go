package main

import "fmt"

var funcVar func(int) int

func incFn(x int) int {
	return x + 1
}

func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

func decFn(x int) int { return x - 1 }

func getMax(vals ...int) int {
	maxV := -1

	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

func main() {
	funcVar = incFn
	fmt.Println(funcVar(1))

	applyIt(incFn, 1)
	applyIt(decFn, 1)

	fmt.Println(getMax(5, 4, 4, 8))

	vslice := []int{1, 2, 4, 9}
	fmt.Println(getMax(vslice...))
}
