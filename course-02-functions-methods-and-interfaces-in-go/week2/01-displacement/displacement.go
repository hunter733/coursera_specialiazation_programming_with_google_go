package main

import (
	"fmt"
	"math"
)

func main() {

	// prompts the user to enter values for acceleration, initial velocity, and initial displacement
	fmt.Printf("Please enter the values of acceleration, initial velocity and initial displacement in separate lines respectively: \n")

	var acceleration float64
	var initialVelocity float64
	var initialDisplacement float64
	var time float64

	n, err := fmt.Scan(&acceleration, &initialVelocity, &initialDisplacement)
	if n == 3 && err != nil {
		panic(err)
	}

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	// prompt the user to enter a value for time
	fmt.Printf("Please enter the value for time: \n")

	ok, err := fmt.Scan(&time)
	if ok == 1 && err != nil {
		panic(err)
	}

	// compute and print the displacement
	fmt.Printf("The displacement after %f seconds is %f\n", time, fn(time))
}

// write a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so and returns a function
func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return ((0.5 * a) * (math.Pow(t, 2))) + vo*t + so
	}
	return fn
}
