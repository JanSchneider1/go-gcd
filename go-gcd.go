package main

import (
	"fmt"
	"go-gcd/gcd"
	"os"
	"strconv"
)

func main() {
	// Get command line arguments
	args := os.Args[1:]

	// Validate argument length
	if len(args) != 2 {
		fmt.Println("please enter command in the following format:")
		fmt.Println("go-gcd <integer> <integer>")
		return
	}

	// Get input
	a, aErr := strconv.Atoi(args[0])
	b, bErr := strconv.Atoi(args[1])

	// Validate input
	if aErr != nil || bErr != nil {
		fmt.Println("please enter command in the following format:")
		fmt.Println("go-gcd <integer> <integer>")
		return
	}

	// Calculate and print result
	result, resultErr := gcd.Gcd(a, b)
	if resultErr != nil {
		fmt.Println(resultErr)
	} else {
		fmt.Println(result)
	}
}
