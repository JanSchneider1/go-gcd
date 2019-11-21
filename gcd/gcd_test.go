package gcd

import (
	"testing"
)

func TestGcdReturnsCorrectResult(t *testing.T) {
	// Given
	cases := []struct{ a, b, expected int }{
		{1, 1, 1},
		{12, 1, 1},
		{12, 3, 3},
		{100, 16, 4},
		{16, 100, 4},
		{258, 18, 6},
		{312, 76, 4},
	}
	// When
	for _, c := range cases {
		actual, _, err := Gcd(c.a, c.b)
		// Then
		if err != nil || actual != c.expected {
			t.Errorf("expected gdc(%d, %d) to be %d, but got %d", c.a, c.b, c.expected, actual)
		}
	}
}

func TestGcdReturnsCorrectSteps(t *testing.T) {
	// Given
	expectedSteps := []Step{
		{100, 51, 49},
		{51, 49, 2},
		{49, 2, 1},
	}
	// When
	_, actualSteps, _ := Gcd(100, 51)
	// Then
	if !Compare(actualSteps, expectedSteps) {
		t.Errorf("actual steps %d do not match expected steps %d", actualSteps, expectedSteps)
	}
}

func TestGcdThrowsErrorForNegativeNumbers(t *testing.T) {
	// Given
	cases := []struct{ a, b int }{
		{-1, -1},
		{12, -1},
		{-1, 12},
	}
	// When
	for _, c := range cases {
		_, _, err := Gcd(c.a, c.b)
		// Then
		if err == nil {
			t.Errorf("expected error to be thrown for gdc(%d, %d)", c.a, c.b)
		}
	}
}

func TestGcdThrowsErrorForZero(t *testing.T) {
	// Given
	cases := []struct{ a, b int }{
		{0, 0},
		{0, 12},
		{12, 0},
	}
	// When
	for _, c := range cases {
		_, _, err := Gcd(c.a, c.b)
		// Then
		if err == nil {
			t.Errorf("expected error to be thrown for gdc(%d, %d)", c.a, c.b)
		}
	}
}
