package gcd

// Saves a single step during calculation
type Step struct {
	A, B, Rest int
}

// Compares two Steps for equality
func (s *Step) Equal(other Step) bool {
	return s.A == other.A && s.B == other.B && s.Rest == other.Rest
}

// Compares two []Step for equality
func Compare(a, b []Step) bool {
	for i, value := range a {
		if !value.Equal(b[i]) {
			return false
		}
	}
	return true
}
