package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Fname string
	Lname string
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var fileName string
	fmt.Println("Please enter the file name")
	fileName, _ = reader.ReadString('\n')
	fileName = strings.Replace(fileName, "\r", "", -1)
	fileName = strings.Replace(fileName, "\n", "", -1)

	readFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var line string
	var people []Person

	for fileScanner.Scan() {
		line = fileScanner.Text()
		stringSlice := strings.Split(line, " ")
		p := Person{Fname: substring(stringSlice[0], 0, 20), Lname: substring(stringSlice[1], 0, 20)}
		people = append(people, p)
	}

	readFile.Close()

	for _, p := range people {
		printPerson(p)
	}
}

func substring(s string, inicial int, lenght int) string {
	if len(s) > lenght {
		return s[inicial : lenght-inicial]
	} else {
		return s
	}
}

func printPerson(p Person) {
	jsonstring, _ := json.Marshal(p)
	fmt.Printf("%s\n", string(jsonstring))
}
