package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	m := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter the name")
	m["name"], _ = reader.ReadString('\n')
	m["name"] = strings.Replace(m["name"], "\r", "", -1)
	m["name"] = strings.Replace(m["name"], "\n", "", -1)

	fmt.Println("Please enter the address")
	m["address"], _ = reader.ReadString('\n')
	m["address"] = strings.Replace(m["address"], "\r", "", -1)
	m["address"] = strings.Replace(m["address"], "\n", "", -1)

	jsonstring, _ := json.Marshal(m)

	fmt.Printf("json: %s", jsonstring)
}
