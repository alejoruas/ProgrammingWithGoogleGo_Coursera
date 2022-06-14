package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var strvalue string
	fmt.Println("Write the text:")

	reader := bufio.NewReader(os.Stdin)

	strvalue, _ = reader.ReadString('\n')
	strvalue = strings.Trim(strvalue, "\n")
	strvalue = strings.Trim(strvalue, "\r")

	if strings.HasPrefix(strings.ToLower(strvalue), "i") &&
		strings.Contains(strings.ToLower(strvalue), "a") &&
		strings.HasSuffix(strings.ToLower(strvalue), "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
