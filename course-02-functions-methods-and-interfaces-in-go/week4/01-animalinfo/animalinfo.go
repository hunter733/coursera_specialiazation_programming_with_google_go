package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Name() string
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name, food, locomotion, noise string
}

type Bird struct {
	name, food, locomotion, noise string
}

type Snake struct {
	name, food, locomotion, noise string
}

func (c Cow) Name() string {
	return c.name
}

func (c Cow) Eat() {
	fmt.Printf(c.food)
}

func (c Cow) Move() {
	fmt.Printf(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Printf(c.noise)
}

func (b Bird) Eat() {
	fmt.Printf(b.food)
}

func (b Bird) Name() string {
	return b.name
}

func (b Bird) Move() {
	fmt.Printf(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Printf(b.noise)
}

func (s Snake) Name() string {
	return s.name
}

func (s Snake) Eat() {
	fmt.Printf(s.food)
}

func (s Snake) Move() {
	fmt.Printf(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Printf(s.noise)
}

func main() {

	fmt.Printf("Available Commands:\n")
	fmt.Printf("newanimal <animal_name> <animal_type> (e.g. newanimal Brian cow)\n")
	fmt.Printf("query <animal_name> <info_required> (e.g. query Brian eat)\n")

	//create an empty struct slice
	var animal = make([]Animal, 0)

	//program should continue in this loop forever
	//program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt.
	for {

		//program should present the user with a prompt, “>”, to indicate that the user can type a request.
		fmt.Printf("\n>")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		words := strings.Fields(input)

		if len(words) >= 3 {

			switch words[0] {
			case "newanimal":

				switch words[2] {
				case "cow":
					c := Cow{name: words[1], food: "grass", locomotion: "walk", noise: "moo"}
					animal = append(animal, c)
				case "bird":
					b := Bird{name: words[1], food: "worms", locomotion: "fly", noise: "peep"}
					animal = append(animal, b)
				case "snake":
					s := Snake{name: words[1], food: "mice", locomotion: "slither", noise: "hsss"}
					animal = append(animal, s)
				}
				fmt.Printf("Created it!")

			case "query":

				for _, a := range animal {
					if a.Name() == words[1] {
						switch words[2] {
						case "eat":
							a.Eat()
						case "move":
							a.Move()
						case "speak":
							a.Speak()
						default:
							fmt.Printf("Don't know about %s %s", words[1], words[2])
						}
					}
				}
			}
		} else {
			fmt.Printf("There was an error parsing the input\n")
		}
	}
}
