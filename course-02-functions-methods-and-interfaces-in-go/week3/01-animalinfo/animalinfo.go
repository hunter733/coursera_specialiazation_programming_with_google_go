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

//The program allows the user to get information about a predefined set of animals.
//Three animals are predefined, cow, bird, and snake.
//Each animal can eat, move, and speak.
//The user can issue a request to find out one of three things about an animal:
//1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks.

func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	fmt.Printf("Please enter two strings <animal> <info_required> (e.g. cow eat)\n")

	//program should continue in this loop forever
	//program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt.
	for {
		//prompts the user to enter two strings
		//request from the user must be a single line containing 2 strings.
		//The first string is the name of an animal, either “cow”, “bird”, or “snake”.
		//The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.

		//program should present the user with a prompt, “>”, to indicate that the user can type a request.
		fmt.Printf("\n>")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		words := strings.Fields(input)

		if len(words) >= 2 {
			animal := Animal{}
			switch words[0] {
			case "cow":
				animal = cow
			case "bird":
				animal = bird
			case "snake":
				animal = snake
			}

			switch words[1] {
			case "eat":
				fmt.Printf(animal.Eat())
			case "move":
				fmt.Printf(animal.Move())
			case "speak":
				fmt.Printf(animal.Speak())
			default:
				fmt.Printf("Don't know about %s %s", words[0], words[1])
			}
		} else {
			fmt.Printf("There was an error parsing the input\n")
		}
	}
}
