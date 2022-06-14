package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func divideSli(sli []int, left, right int) []int {

	return sli[left:right]
}

func sortSli(sli []int, c chan []int) {
	sort.Ints(sli)

	c <- sli
}

func printFun(sli []int) {
	fmt.Print("Sorted subarray: ")
	fmt.Println(sli)
}

func merge(s1, s2 []int) []int {
	lenS1 := len(s1)
	lenS2 := len(s2)
	newSli := make([]int, lenS1+lenS2)
	var i, j, idx int
	for i < lenS1 && j < lenS2 {
		if s1[i] <= s2[j] {
			newSli[idx] = s1[i]
			idx++
			i++
		} else {
			newSli[idx] = s2[j]
			idx++
			j++
		}
	}
	for i < lenS1 {
		newSli[idx] = s1[i]
		idx++
		i++
	}
	for j < lenS2 {
		newSli[idx] = s2[j]
		idx++
		j++
	}

	return newSli
}

func main() {

	fmt.Print("Please enter a series of number with space: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numberLine := scanner.Text()
	numbers := strings.Split(string(numberLine), " ")
	numberSli := make([]int, len(numbers))
	for i := range numberSli {
		numberSli[i], _ = strconv.Atoi(numbers[i])
	}
	n := len(numberSli) / 4
	sli1 := divideSli(numberSli, 0, n)
	sli2 := divideSli(numberSli, n, 2*n)
	sli3 := divideSli(numberSli, 2*n, 3*n)
	sli4 := divideSli(numberSli, 3*n, len(numberSli))
	//fmt.Println(sli1, sli2, sli3, sli4, numberSli)
	c := make(chan []int, 4)
	go sortSli(sli1, c)
	go sortSli(sli2, c)
	go sortSli(sli3, c)
	go sortSli(sli4, c)

	sortedS1 := <-c
	sortedS2 := <-c
	sortedS3 := <-c
	sortedS4 := <-c
	printFun(sortedS1)
	printFun(sortedS2)
	printFun(sortedS3)
	printFun(sortedS4)

	time.Sleep(100 * time.Millisecond)
	fmt.Print("Sorted array: ")
	M1 := merge(sortedS1, sortedS2)
	M2 := merge(sortedS3, sortedS4)
	finalSli := merge(M1, M2)
	fmt.Print(finalSli)
}
