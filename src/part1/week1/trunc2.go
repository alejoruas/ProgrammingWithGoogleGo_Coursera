package main

import (
	"fmt"
	"math"
)

func main() {
	var floater float64

	fmt.Println("Enter a floating point number:")
	fmt.Scanf("%f", &floater)

	i := math.Round(floater)
	fmt.Print(i)
}