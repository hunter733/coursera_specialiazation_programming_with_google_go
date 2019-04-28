package main

import "fmt"

func main() {
	var floatNum float32

	fmt.Printf("Please enter a Floating Point Number: ")
	ok, err :=
		fmt.Scan(&floatNum)

	if ok == 1 {
		fmt.Printf("The truncated input is: %d\n", int(floatNum))
	} else {
		fmt.Printf("There was an error in the input: %v\n", err.Error())
	}
}
