package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

func (c Cow) Eat() { fmt.Println("grass") }

func (c Cow) Move() { fmt.Println("walk") }

func (c Cow) Speak() { fmt.Println("moo") }

type Bird struct {
	name string
}

func (b Bird) Eat() { fmt.Println("worms") }

func (b Bird) Move() { fmt.Println("fly") }

func (b Bird) Speak() { fmt.Println("peep") }

type Snake struct {
	name string
}

func (s Snake) Eat() { fmt.Println("mice") }

func (s Snake) Move() { fmt.Println("slither") }

func (s Snake) Speak() { fmt.Println("hsss") }

func main() {
	fmt.Println("We have 3 type of animals (cow, bird, snake) that can eat, move or speak")
	fmt.Println("Type a request... (newanimal, query or type exit to finish)")

	instruccion := []string{"", ""}

	animals := make(map[string]Animal)

	fmt.Print(">")
	instruccion = ReadInstruccion()

	for instruccion[0] != "exit" {

		switch instruccion[0] {
		case "newanimal":
			newAnimal(animals, instruccion[2], instruccion[1])
		case "query":
			request(animals, instruccion[1], instruccion[2])
		default:
			fmt.Print("unknown command\n")
		}

		fmt.Print(">")
		instruccion = ReadInstruccion()
	}
}

func request(animals map[string]Animal, name string, info string) {
	var animal Animal
	animal = animals[name]

	if animal == nil {
		fmt.Print("the animal does not exist\n")
		return
	}

	switch info {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Print("unknown command\n")
	}
}

func newAnimal(animals map[string]Animal, animalType string, name string) {
	var newAnimal Animal

	switch animalType {
	case "cow":
		newAnimal = Cow{name: name}
	case "bird":
		newAnimal = Bird{name: name}
	case "snake":
		newAnimal = Snake{name: name}
	default:
		{
			fmt.Print("unknown animal type\n")
			return
		}
	}
	animals[name] = newAnimal
}

func ReadInstruccion() []string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.Replace(line, "\r", "", -1)
	line = strings.Replace(line, "\n", "", -1)

	returnValue := strings.Split(line, " ")

	return returnValue
}
