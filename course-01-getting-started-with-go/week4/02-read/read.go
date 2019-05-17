package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {

	//create an empty struct slice
	var sli = make([]Person, 0, 10)

	// PROMPT: Your program should prompt the user for the name of the text file.
	// prompts the user to enter file with extension
	var file string

	fmt.Printf("Please enter file (with ext) to read: ")

	ok, err :=
		fmt.Scan(&file)

	if ok == 1 && err == nil {

		file, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// PROMPT: Your program will successively read each line of the text file
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			var line = scanner.Text()
			// PROMPT: Your program will create a struct which contains the first and last names found in the file.
			// PROMPT: Each struct created will be added to a slice
			// PROMPT: After all lines have been read from the file, your program will have a slice containing one struct for each line in the file
			sli = addToSlice(line, sli)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		// PROMPT: your program should iterate through your slice of structs and print the first and last names found in each struct.
		for _, v := range sli {
			fmt.Printf(v.fname + " " + v.lname + "\n")
		}
	}
}

func addToSlice(fullName string, sli []Person) []Person {
	if fullName != "" {
		// splits the first name and last name with space
		splittedFullName := strings.Fields(fullName)

		// declares a struct instance and assigns the values
		var p Person

		if len(splittedFullName) > 0 {
			p.fname = splittedFullName[0]
		}

		if len(splittedFullName) > 1 {
			p.lname = splittedFullName[1]
		}

		// adds the struct instance to slice
		sli = append(sli, p)
	}

	return sli
}
