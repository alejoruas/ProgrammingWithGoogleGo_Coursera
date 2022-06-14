package main

import (
	"fmt"
	"strconv"
)

func main() {
	var option string
	var floatValue float64 = 0
	var strvalue string

	for {
		fmt.Println("Write a float number")
		fmt.Scan(&strvalue)
		floatValue, _ = strconv.ParseFloat(strvalue, 64)
		fmt.Printf("The truncate number is %d\n", int32(floatValue))

		fmt.Printf("Write 'S' to continue or something else to exit\n")
		fmt.Scan(&option)

		if option != "S" {
			break
		}
	}
}
