package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) InitMe(food, locomotion, noise string) {
	a.food = food
	a.locomotion = locomotion
	a.noise = noise
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {

	fmt.Println("We have 3 animals (cow, bird, snake) that can eat, move or speak")
	fmt.Println("Type a request... (type exit to finish)")

	instruccion := []string{"", ""}

	var cow Animal
	var bird Animal
	var snake Animal

	cow.InitMe("grass", "walk", "moo")
	bird.InitMe("worms", "fly", "peep")
	snake.InitMe("mice", "slither", "hsss")

	fmt.Print("\n>")
	instruccion = ReadInstruccion()

	for instruccion[0] != "exit" {
		if instruccion[0] == "cow" {
			if instruccion[1] == "eat" {
				cow.Eat()
			}
			if instruccion[1] == "move" {
				cow.Move()
			}
			if instruccion[1] == "speak" {
				cow.Speak()
			}
		}

		if instruccion[0] == "bird" {
			if instruccion[1] == "eat" {
				bird.Eat()
			}
			if instruccion[1] == "move" {
				bird.Move()
			}
			if instruccion[1] == "speak" {
				bird.Speak()
			}
		}

		if instruccion[0] == "snake" {
			if instruccion[1] == "eat" {
				snake.Eat()
			}
			if instruccion[1] == "move" {
				snake.Move()
			}
			if instruccion[1] == "speak" {
				snake.Speak()
			}
		}

		fmt.Print("\n>")
		instruccion = ReadInstruccion()
	}
}

func ReadInstruccion() []string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.Replace(line, "\r", "", -1)
	line = strings.Replace(line, "\n", "", -1)

	returnValue := []string{"", ""}
	returnValue = strings.Split(line, " ")

	return returnValue
}
