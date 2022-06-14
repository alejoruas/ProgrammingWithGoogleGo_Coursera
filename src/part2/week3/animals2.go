package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

type Animal struct{
  eat string
  move string
  speak string
}

func (animal Animal) Eat() string { return animal.eat }
func (animal Animal) Move() string { return animal.move }
func (animal Animal) Speak() string { return animal.speak }

func main() {
  animals := map[string]Animal{
    "cow":   Animal{ eat: "grass", move: "walk",    speak: "moo" },
    "bird":  Animal{ eat: "worms", move: "fly",     speak: "peep" },
    "snake": Animal{ eat: "mice",  move: "slither", speak: "hsss" },
  }

  for true {
    command := ReadCommand()
    result := ProcessCommand(animals, command)
    fmt.Println(result)
  }
}

func ProcessCommand(animals map[string]Animal, command map[string]string) string {
  animal := animals[command["animal"]]

  switch command["action"] {
  case "eat":
    return animal.Eat()
  case "move":
    return animal.Move()
  case "speak":
    return animal.Speak()
  default:
    return "unknown action"
  }
}

func ReadCommand() map[string]string {
  var animal, action, errorMessage string

  line := ReadLine()
  animal, action, errorMessage = ValidateLine(line)
  for errorMessage != "" {
    fmt.Println("format: <anmial> <action> ", errorMessage)
    line := ReadLine()
    animal, action, errorMessage = ValidateLine(line)
  }

  return map[string]string{ "animal": animal, "action": action }
}

func ReadLine() string {
  scanner := bufio.NewScanner(os.Stdin)
  userInput := ""

  fmt.Printf("> ")
  if scanner.Scan() { userInput = scanner.Text() }

  return userInput
}

func ValidateLine(line string) (string, string, string) {
  validAnimals := []string{"cow", "bird", "snake"}
  validActions := []string{"eat", "move", "speak"}

  words := strings.Fields(line)
  if len(words) != 2 {
    return "", "", "more than two words detected."
  }
  if !InArray(words[0], validAnimals) {
    return "", "", "not a valid animal, must be one of these: " + strings.Join(validAnimals, ", ")
  }
  if !InArray(words[1], validActions) {
    return "", "", "not a valid action, must be one of these: " + strings.Join(validActions, ", ")
  }

  return words[0], words[1], ""
}

func InArray(data string, dataArray []string) bool {
  for _, v := range dataArray {
		if v == data {
			return true
		}
	}

	return false
}
