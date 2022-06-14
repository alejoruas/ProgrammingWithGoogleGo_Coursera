package main

import (
	"fmt"
	"strconv"
)

func main() {
	var text string
	sli := make([]int, 0, 3)

	for {
		fmt.Println("Please enter a integer number o X to exit")
		fmt.Scan(&text)

		if text == "X" {
			break
		}

		valNew, _ := strconv.Atoi(text)
		insertInOrder(valNew, &sli)

		showArray("", &sli)
	}

}

func insertInOrder(valNew int, array *[]int) {
	inserted := false

	for i, val := range *array {
		if val > valNew {
			inserted = true
			*array = append(*array, 0)

			for j := len(*array) - 1; j > i; j-- {
				(*array)[j] = (*array)[j-1]
			}

			(*array)[i] = valNew
			break
		}
	}

	if !inserted {
		*array = append(*array, valNew)
	}
}

func showArray(name string, array *[]int) {
	fmt.Printf("Array %s\n", name)
	for i, val := range *array {
		fmt.Printf("Index: %d value: %d\n", i, val)
	}
	fmt.Printf("End Array %s\n", name)
}
