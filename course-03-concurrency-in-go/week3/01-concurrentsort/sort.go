package main

import (
	"fmt"
	"strconv"
)

var partitioned = make([][]int, 0)

// PROMPT: Write a program to sort an array of integers.
func main() {
	//create an empty integer slice of dynamic size
	var sli = make([]int, 0)

	//PROMPT: The program should prompt the user to input a series of integers.
	fmt.Printf("Please enter integers (one per line) or X to exit: \n")
	for {
		//prompts the user to enter integers
		var input string

		ok, err :=
			fmt.Scan(&input)

		if ok == 1 && err == nil {
			//The program quits (exiting the loop) when the user enters the character ‘X’ instead of an integer.
			if input == "X" {
				c1 := make(chan []int)
				c2 := make(chan []int)
				c3 := make(chan []int)
				c4 := make(chan []int)

				//PROMPT: The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
				//PROMPT: Each partition should be of approximately equal size.
				Partition(sli, 4)

				//PROMPT: The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
				for i := 0; i < len(partitioned); i++ {
					go BubbleSort(partitioned[i], c1)
				}
				a := <-c1
				b := <-c1
				c := <-c1
				d := <-c1

				//PROMPT: Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
				fmt.Println("Routine#1:", a)
				fmt.Println("Routine#2:", b)
				fmt.Println("Routine#3:", c)
				fmt.Println("Routine#4:", d)

				//PROMPT: Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.
				var sli2 = make([]int, 0)
				sli2 = append(sli2, a...)
				sli2 = append(sli2, b...)

				var sli3 = make([]int, 0)
				sli3 = append(sli3, c...)
				sli3 = append(sli3, d...)

				go BubbleSort(sli2, c2)
				go BubbleSort(sli3, c3)

				e := <-c2
				f := <-c3

				var sli4 = make([]int, 0)
				sli4 = append(sli4, e...)
				sli4 = append(sli4, f...)

				go BubbleSort(sli4, c4)
				g := <-c4

				//PROMPT: When sorting is complete, the main goroutine should print the entire sorted list.
				fmt.Println("Sorted Array: ")
				fmt.Println(g)

				break
			}

			integerNum, err := strconv.Atoi(input)

			if err == nil {
				//stores the integers in a slice
				sli = append(sli, integerNum)
			}

		}
	}

}

// write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing.
// The BubbleSort() function should modify the slice so that the elements are in sorted order.
func BubbleSort(sli []int, c chan []int) {
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
	c <- sli
}

//write a Swap() function which performs this operation.
//Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice.
//The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.
func Swap(sli []int, i int) {
	var temp = sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = temp
}

func Partition(sli []int, routines int) {

	chunkSize := (len(sli) + routines - 1) / routines

	for i := 0; i < len(sli); i += chunkSize {
		end := i + chunkSize

		if end > len(sli) {
			end = len(sli)
		}

		partitioned = append(partitioned, sli[i:end])
	}

	for len(partitioned) < 4 {
		fmt.Printf("Oops! not enough elements to partition equally. Adjusting...\n")
		partitioned = append(partitioned, []int{})
	}

	// fmt.Printf("%#v\n", partitioned)
}
