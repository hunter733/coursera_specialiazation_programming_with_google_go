/*
Race condition is read-write access to one variable in 2 or more different concur routines.

I've wrote some example here.

GoRoutine switches GotMain to true only if it is false.
But because race condition one goroutine switch GotMain to true when GotMain is already true.
Every program run will print not determining output.
For example:

I am 2. I am leader.
I am 2. Switch GotMain to true, it is false now
I am 1. I am leader.
I am 1. Switch GotMain to true, it is true now


I am 1. I am leader.
I am 1. Switch GotMain to true, it is false now
I am 2. I am leader.
I am 2. Switch GotMain to true, it is true now


I am 2. I am leader.
I am 1. I am leader.
I am 1. Switch GotMain to true, it is false now
I am 2. Switch GotMain to true, it is false now
*/

package main

import (
	"fmt"
	"time"
)

var GotMain = false

func GoRoutine (index int) {
	msg := "I am %d. I am %s.\n"
	if GotMain == true {
		fmt.Printf(msg, index, "follower")
	} else {
		time.Sleep(1 * time.Millisecond)         // need to switch to another goroutine
		fmt.Printf(msg, index, "leader")
		fmt.Printf("I am %d. Switch GotMain to true, it is %t now\n", index, GotMain)
		GotMain = true
	}
}

func main() {
	go GoRoutine(1)
	go GoRoutine(2)
	time.Sleep(10 * time.Millisecond)
}
