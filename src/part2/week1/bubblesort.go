package main

import "fmt"

func main() {
	array := make([]int, 10, 10)

	fmt.Println("Please type 10 integers")

	for i := 0; i < 10; i++ {
		var value int
		fmt.Scan(&value)
		array[i] = value
	}

	fmt.Println("Initial Array: ")
	showArray(array)
	bublesort(array)
	fmt.Println("\nSorted Array: ")
	showArray(array)
}

func showArray(array []int) {
	for _, value := range array {
		fmt.Printf("%d - ", value)
	}
}

func bublesort(array []int) {
	i := 0
	sorted := false

	for i < len(array) && sorted == false {
		i = i + 1
		sorted = true

		for j := 0; j < len(array)-i; j++ {
			if (array)[j] > (array)[j+1] {
				sorted = false
				swap(array, j)
			}
		}
	}
}

func swap(array []int, index int) {
	aux := (array)[index]
	(array)[index] = (array)[index+1]
	(array)[index+1] = aux
}
