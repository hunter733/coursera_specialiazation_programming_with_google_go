package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

const NewAnimal = "newanimal"
const Query = "query"

func main() {

	consoleReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf(">")
		input, err := consoleReader.ReadString('\n')
		_ = err

		stringSlice := strings.Split(strings.TrimSpace(input), " ")

		firstArgument := func() string {
			if len(stringSlice) > 0 {
				return strings.ToLower(stringSlice[0])
			} else {
				return ""
			}
		}()

		secondArgument := func() string {
			if len(stringSlice) > 1 {
				return strings.Title(stringSlice[1])
			} else {
				return ""
			}
		}()

		thirdArgument := func() string {
			if len(stringSlice) > 2 {
				return strings.Title(stringSlice[2])
			} else {
				return ""
			}
		}()

		if (firstArgument == NewAnimal) {
			animal := GetAnimalByType(thirdArgument)
			if animal!=nil {
				fmt.Println("Created it!")
			}

		}

		if (firstArgument == Query) {
			animal := GetAnimalByType(secondArgument)

			if animal != nil {
				f := reflect.ValueOf(animal).MethodByName(thirdArgument)
				if f.IsValid() {
					f.Call([]reflect.Value{})
				} else {
					fmt.Println("Unknown action")
				}
			}

		}
	}

}

func GetAnimalByType(animalType string) Animal {
	switch animalType {
	case "Cow":
		return Cow{"grass","walk","moo"}
		break
	case "Snake":
		return Snake{"mice","slither","hsss"}
		break
	case "Bird":
		return Bird{"worm","fly","peep"}
		break
	default:
		fmt.Println("Unknown animal")
	}

	return nil
}

type Animal interface {
	Speak()
	Eat()
	Move()
}
type Cow struct {
	food       string
	locomotion string
	noise      string
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}
