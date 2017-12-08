package solutions

import (
	"testing"
)

// Multiplesof3and5 ...
func Multiplesof3and5(number int) int {
	result := 0
	for i := 1; i < number; i++ {

		b := false
		b = b || (i%3 == 0)
		b = b || (i%5 == 0)
		if b {
			result += i
		}
	}
	return result
}

// https://projecteuler.net/problem=1
func TestSolution01(t *testing.T) {
	ans := Multiplesof3and5(1000)
	if ans != 233168 {
		t.Fatal("")
	}
}
