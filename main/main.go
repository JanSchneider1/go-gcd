package main

import (
	"fmt"
	"go-gcd"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	a, _ := strconv.Atoi(args[0])
	b, _ := strconv.Atoi(args[1])
	result, _ := gcd.Gcd(a, b)
	fmt.Printf(string(result))
}
