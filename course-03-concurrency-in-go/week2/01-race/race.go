package main

import (
	"fmt"
	"time"
)

var i int

// The program contains two go routines demonstrating data race
// 1) The value of i is being set to 5
// 2) The value of i is being returned from the function
// Now depending on which of these two routines complete first, our value printed will be either 0 (the default integer value) or 5.

func main() {
	go setNumber()
	// time.Sleep(100 * time.Millisecond) //delay until the setNumber routine assigns the number
	go fmt.Println(getNumber())
	time.Sleep(100 * time.Millisecond) //delay until the getNumber routine prints the number
}

func getNumber() int {
	return i
}

func setNumber() {
	i = 5
}
