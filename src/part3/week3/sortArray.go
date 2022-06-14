package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	array := ReadArray()

	var count int
	var indexBegin, indexEnd int
	count = len(array) / 4

	indexBegin = 0
	indexEnd = count
	array1 := array[indexBegin:indexEnd]

	indexBegin = indexEnd
	indexEnd = indexBegin + count
	array2 := array[indexBegin:indexEnd]

	indexBegin = indexEnd
	indexEnd = indexBegin + count
	array3 := array[indexBegin:indexEnd]

	indexBegin = indexEnd
	indexEnd = len(array)
	array4 := array[indexBegin:indexEnd]

	c := make(chan []int, 4)
	go bublesortThread(array1, c)
	go bublesortThread(array2, c)
	go bublesortThread(array3, c)
	go bublesortThread(array4, c)

	array1 = <-c
	array2 = <-c
	array3 = <-c
	array4 = <-c

	newArray := append(array1, array2...)
	newArray = append(newArray, array3...)
	newArray = append(newArray, array4...)

	newArray = bublesort(newArray)
	showArray("FINISH", &newArray)
}

func ReadArray() []int {
	array := make([]int, 0, 0)

	var text string

	for {
		fmt.Println("Please enter a integer number o X to finish")
		fmt.Scan(&text)

		if strings.ToLower(text) == "x" {
			break
		}

		valNew, _ := strconv.Atoi(text)
		array = append(array, valNew)
	}

	showArray("Complete array to sort:", &array)

	return array
}

func ToLower(text string) {
	panic("unimplemented")
}

func showArray(name string, array *[]int) {
	fmt.Printf("%s\n", name)
	for i, val := range *array {
		fmt.Printf("Index: %d value: %d\n", i, val)
	}
	fmt.Printf("\n")
}

func bublesortThread(array []int, c chan []int) {
	c <- bublesort(array)
}

func bublesort(array []int) []int {
	i := 0
	sorted := false

	showArray("subarray to sort", &array)

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
	return array
}

func swap(array []int, index int) {
	aux := (array)[index]
	(array)[index] = (array)[index+1]
	(array)[index+1] = aux
}
