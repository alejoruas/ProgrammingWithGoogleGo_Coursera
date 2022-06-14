package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Person struct {
	Name  string `json:"name"`
	Addr  string `json:"addr"`
	Phone string `json:"phone"`
}

func main() {
	dat, _ := ioutil.ReadFile("person.json")

	var p Person
	json.Unmarshal(dat, &p)

	printPe4rson(p)

	p.Name = "Nuevo nombre"
	p.Addr = "Nueva dir"
	p.Phone = time.Now().String()

	jsonstring, _ := json.Marshal(p)
	ioutil.WriteFile("person.json", jsonstring, 0777)

}

func printPe4rson(p Person) {
	fmt.Printf("\nname:%s", p.Name)
	fmt.Printf("\naddr:%s", p.Addr)
	fmt.Printf("\nphone:%s", p.Phone)
}
