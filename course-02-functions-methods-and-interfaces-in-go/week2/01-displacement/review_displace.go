package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type displacement struct {
	accel float64
	ivel  float64
	idisp float64
}

type dispfunc func(time float64) float64

func GenDisplaceFunc(accel, ivel, idisp float64) dispfunc {
	return func(t float64) float64 {
		return (.5 * accel * math.Pow(t, 2)) + (ivel * t) + idisp
	}
}

func getUserInput(msg string) []string {
	var line string
	fmt.Printf("%s: ", msg)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		line = scanner.Text()
		break
	}

	return strings.Fields(line)

}
func getFloat(paramName string) float64 {
	var err error
	param := -1.0
	for param < 0 {

		inputs := getUserInput("Please enter " + paramName)
		// use only first user input
		if len(inputs) > 0 {
			if param, err = strconv.ParseFloat(inputs[0], 64); err != nil {
				fmt.Printf("Invalid %s\n", paramName)
				param = -1
			}
		}

	}
	return param
}

func getDispacParams() displacement {

	var displace displacement
	displace.accel = getFloat("acceleration")
	displace.ivel = getFloat("initial velosity")
	displace.idisp = getFloat("initial displacement")
	return displace
}

func getTimes() []float64 {
	tStrs := getUserInput("Enter one or more times separated by space")
	times := []float64{}
	for _, time := range tStrs {
		if t, err := strconv.ParseFloat(time, 64); err == nil {
			times = append(times, t)
		}
	}
	return times
}

func main() {
	displace := getDispacParams()
	// to comply with instructions
	dispf := GenDisplaceFunc(displace.accel, displace.ivel, displace.idisp)

	times := getTimes()

	fmt.Printf("Displacements for %d times:\n", len(times))
	for _, time := range times {
		fmt.Printf("  Displacement after %.1f seconds = %.2f\n", time, dispf(time))
	}
}
