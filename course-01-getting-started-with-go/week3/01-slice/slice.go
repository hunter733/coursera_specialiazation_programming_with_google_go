package main

import (
	"fmt"
	"sort"
	"strconv"
)

//The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order
//The slice must grow in size to accommodate any number of integers which the user decides to enter

func main() {
	//create an empty integer slice of size (length) 3
	var sli = make([]int, 0, 3)

	//The program should be written as a loop.
	for {
		//prompts the user to enter integers
		var input string

		fmt.Printf("Please enter an integer: ")

		ok, err :=
			fmt.Scan(&input)

		if ok == 1 {
			//The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
			if input == "X" {
				break
			}

			integerNum, err := strconv.Atoi(input)

			if err == nil {
				//stores the integers in a sorted slice
				sli = append(sli, integerNum)

				sort.Ints(sli)
				fmt.Println(sli)
			} else {
				fmt.Printf("There was an error parsing the input\n")
				break
			}

		} else {
			fmt.Printf("There was an error in the input: %v\n", err.Error())
			break
		}
	}

}
