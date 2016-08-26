package solutions

import (
	"testing"
)
// https://projecteuler.net/problem=1
func TestSolution01(t *testing.T) {
	ans := Multiplesof3and5(1000)
	if ans != 233168 {
		t.Fatal("")
	}
}
