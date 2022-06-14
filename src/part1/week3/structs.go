package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Addr  string `json:"addr"`
	Phone string `json:"phone"`
}

func main() {
	p1 := new(Person)
	p1.Name = "Alejandro"

	p2 := Person{Name: "Alejo", Addr: "calle", Phone: "235345"}

	fmt.Printf(p2.Name + "\n")

	jsonstring, err := json.Marshal(p2)
	fmt.Print(string(jsonstring))
	fmt.Println()
	fmt.Print(err)

	strp3 := `{"name": "Alejandro", "addr": "dir", "phone": "564564"}`

	var p3 Person
	json.Unmarshal([]byte(strp3), &p3)
	fmt.Println("Person: ", p3)
	printPe4rson(p3)

	jsonstring, err = json.Marshal(p3)
	fmt.Print(string(jsonstring))
	fmt.Println()
	fmt.Print(err)

}

func printPe4rson(p Person) {
	fmt.Printf("\nname:%s", p.Name)
	fmt.Printf("\naddr:%s", p.Addr)
	fmt.Printf("\nphone:%s", p.Phone)
}
