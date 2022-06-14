package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sequence := getSequence()
	BubbleSort(sequence)
	printSequence(sequence)
}

func getSequence() []int{
	input := getInputString()
	return convertStringToIntSlice(input)
}

func getInputString() string {
	fmt.Print("Enter a sequence of up to 10 integers: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func convertStringToIntSlice(intString string) []int {
	var integers []int
	for _, field := range strings.Split(intString, " ") {
		value, _ := strconv.Atoi(field)
		integers = append(integers, value)
	}
	return integers
}

func BubbleSort(sequence []int) {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(sequence)-1; i++ {
			if sequence[i] > sequence[i+1] {
				sorted = false
				Swap(sequence, i)
			} 
		}
	}
}

func Swap(sequence []int, index int) {
	current := sequence[index]
	sequence[index] = sequence[index+1]
	sequence[index+1] = current
}

func printSequence(sequence []int) {
	fmt.Println(strings.Trim(fmt.Sprint(sequence), "[]"))
}