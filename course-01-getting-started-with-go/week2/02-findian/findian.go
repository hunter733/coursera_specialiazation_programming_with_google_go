package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string

	fmt.Printf("Please enter a string: ")

	reader := bufio.NewReader(os.Stdin)
	bytes, _, err := reader.ReadLine()
	input = string(bytes)

	if err == nil {
		var lowerCaseInput = strings.ToLower(input)
		var strLength = len(lowerCaseInput)
		var foundCharA = false

		if strLength < 3 { //check if at all the input string has the possibility of having i,a,n
			fmt.Printf("Not Found!\n")
			return
		}

		var strBetweenEdges = lowerCaseInput[1 : len(lowerCaseInput)-1]
		foundCharA = strings.Contains(strBetweenEdges, "a") //check if string edges contains character 'a'

		if foundCharA == false {
			fmt.Printf("Not Found!\n")
			return
		}

		//check if string starts with character 'i'
		if lowerCaseInput[0] != 'i' {
			fmt.Printf("Not Found!\n")
			return
		}

		//check if string ends with character 'n'
		if lowerCaseInput[strLength-1] != 'n' {
			fmt.Printf("Not Found!\n")
			return
		}

		fmt.Printf("Found!\n")

	} else {
		fmt.Printf("There was an error in the input: %v\n", err.Error())
	}
}
