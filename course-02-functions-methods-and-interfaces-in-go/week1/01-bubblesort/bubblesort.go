package main

import (
	"fmt"
	"strconv"
)

func main() {
	//create an empty integer slice of size (length) 10
	var sli = make([]int, 0, 10)

	var maxInputCount = 10
	var inputCount = 0

	// The program should prompt the user to type in a sequence of up to 10 integers.
	// The program should print the integers out on one line, in sorted order, from least to greatest.

	fmt.Printf("Please enter upto 10 integers (one per line) or X to exit: \n")
	for {
		//prompts the user to enter integers
		var input string

		ok, err :=
			fmt.Scan(&input)

		if ok == 1 && err == nil {
			//The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
			if input == "X" {
				BubbleSort(sli)
				fmt.Println(sli)
				break
			}

			integerNum, err := strconv.Atoi(input)

			if err == nil {
				inputCount++

				//stores the integers in a sorted slice
				sli = append(sli, integerNum)
				if inputCount == maxInputCount {
					BubbleSort(sli)
					fmt.Println(sli)
					break
				}

			}

		}
	}

}

// write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing.
// The BubbleSort() function should modify the slice so that the elements are in sorted order.
func BubbleSort(sli []int) {
	var isSorted = false
	for !isSorted {
		isSorted = true
		for i := 0; i < len(sli)-1; i++ {
			if sli[i] > sli[i+1] {
				Swap(sli, i)
				isSorted = false
			}
		}
	}
}

//write a Swap() function which performs this operation.
//Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice.
//The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.
func Swap(sli []int, i int) {
	var temp = sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = temp
}
