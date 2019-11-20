package gcd

import "errors"

func gcd(a, b int) (int, error) {
	return greatestCommonDivisor(a, b)
}

func greatestCommonDivisor(a, b int) (int, error) {
	// Throw error if input is 0
	if a == 0 || b == 0 {
		return 0, errors.New("cannot determine gcd with number 0")
	}
	// Throw error if input is negative
	if a < 0 || b < 0 {
		return 0, errors.New("cannot determine gcd for negative numbers")
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
	return b, nil
}
