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

func (a Animal) Eat() string {
	return a.food
}
func (a Animal) Move() string {
	return a.locomotion
}
func (a Animal) Speak() string {
	return a.noise
}
func main() {
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		s := scanner.Text()
		strText := strings.Split(s, " ")
		animal := new(Animal)
		if strText[0] == "cow" {
			animal.food = "grass"
			animal.locomotion = "walk"
			animal.noise = "moo"
		} else if strText[0] == "bird" {
			animal.food = "worms"
			animal.locomotion = "fly"
			animal.noise = "peep"
		} else if strText[0] == "snake" {
			animal.food = "mice"
			animal.locomotion = "slither"
			animal.noise = "hiss"
		} else {
			fmt.Println("Error: Invalid string!")
			break
		}

		if len(strText) < 2 {
			fmt.Println("Eat : ", animal.Eat())
			fmt.Println("Movement : ", animal.Move())
			fmt.Println("Speak : ", animal.Speak())
			break
		}

		if strText[1] == "eat" {
			fmt.Println(animal.Eat())
		} else if strText[1] == "move" {
			fmt.Println(animal.Move())
		} else if strText[1] == "speak" {
			fmt.Println(animal.Speak())
		} else {
			fmt.Println("Error: Invalid string!")
			break
		}
	}
}
