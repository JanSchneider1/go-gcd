package gcd

import "errors"

func Gcd(a, b int) (int, []Step, error) {
	return greatestCommonDivisor(a, b)
}

func greatestCommonDivisor(a, b int) (int, []Step, error) {
	// Throw error if input is 0
	if a == 0 || b == 0 {
		return 0, nil, errors.New("cannot determine gcd with number 0")
	}
	// Throw error if input is negative
	if a < 0 || b < 0 {
		return 0, nil, errors.New("cannot determine gcd for negative numbers")
	}
	// Ensure A is bigger that B
	if b > a {
		a, b = b, a
	}
	// Divide until Rest is 0
	rest := a % b
	steps := []Step{{a, b, rest}}
	for i := 0; rest != 0; i++ {
		// Save each step during calculation
		if i != 0 {
			steps = append(steps, Step{a, b, rest})
		}
		// Calculate next step
		a, b = b, rest
		rest = a % b
	}
	return b, steps, nil
}
