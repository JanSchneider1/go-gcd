package main

import "fmt"

func main(){
	fmt.Println(gcd(16, 2))
	fmt.Println(gcd(16, 1))
	fmt.Println(gcd(16, 6))
	fmt.Println(gcd(63, 20))
	fmt.Println(gcd(2530, 120))
}

func gcd(a, b int) int{
	return greatestCommonDivisor(a, b)
}

func greatestCommonDivisor(a, b int) int {
	// Invalid for negative numbers
	if a < 0 || b < 0 {
		return 0
	}
	// Ensure a is bigger that b
	if b > a {
		a, b = b, a
	}
	// Divide until rest is 0
	rest := a % b
	for rest != 0 {
		a, b = b, rest
		rest = a % b
	}
	return b
}
