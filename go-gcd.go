/*
 * Command line tool for calculating the 'greatest common divisor'
 */
package main

import (
	"flag"
	"fmt"
	"go-gcd/gcd"
	"strconv"
)

func main() {
	// Define and parse command line options
	printStepsEnabled := flag.Bool("s", false, "prints each step during calculation")
	flag.Parse()

	// Get command line arguments after parsing options
	args := flag.Args()

	// Validate argument length
	if len(args) != 2 {
		printError()
		return
	}

	// Get input as int
	a, aErr := strconv.Atoi(args[0])
	b, bErr := strconv.Atoi(args[1])

	// Validate input
	if aErr != nil || bErr != nil {
		printError()
		return
	}

	// Calculate and produce output
	result, steps, resultErr := gcd.Gcd(a, b)
	if resultErr != nil {
		fmt.Println(resultErr)
	} else {
		if *printStepsEnabled {
			printSteps(steps)
		} else {
			fmt.Println(result)
		}
	}
}

// Prints each step to console
func printSteps(steps []gcd.Step) {
	for _, value := range steps {
		fmt.Printf("%-5d mod %5d = %5d \n", value.A, value.B, value.Rest)
	}
}

// Prints error message
func printError() {
	fmt.Println("please enter command in the following format:")
	fmt.Println("go-gcd <integer> <integer>")
}
