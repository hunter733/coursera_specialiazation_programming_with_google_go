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

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}
func (c Bird) Eat() {
	fmt.Println("worms")
}
func (c Bird) Move() {
	fmt.Println("fly")
}
func (c Bird) Speak() {
	fmt.Println("peep")
}
func (c Snake) Eat() {
	fmt.Println("mice")
}
func (c Snake) Move() {
	fmt.Println("slither")
}
func (c Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	animals := map[string]Animal{}
	fmt.Print("> ")
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")

		if len(input) != 3 {
			fmt.Println("Invalid number of arguments")
			fmt.Print("> ")
			continue
		}

		action := input[0]
		switch action {
		case "newanimal":
			animal := input[2]
			switch animal {
			case "cow":
				animals[input[1]] = Cow{}
				fmt.Println("Created it!")
			case "bird":
				animals[input[1]] = Bird{}
				fmt.Println("Created it!")
			case "snake":
				animals[input[1]] = Snake{}
				fmt.Println("Created it!")
			default:
				fmt.Println("Invalid animal!")
			}

		case "query":
			animal, ok := animals[input[1]]
			if ok {
				switch input[2] {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				default:
					fmt.Println("Invalid action for animal")
				}
			} else {
				fmt.Println("Animal not found!")
			}

		default:
			fmt.Println("Invalid action. It should be `newanimal` or `query`.")
		}

		fmt.Print("> ")
	}
}
