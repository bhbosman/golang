package sol005

import "testing"

func TestSolution05(t *testing.T) {

	n := MinumumDivisibleNumber([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})

	if n != 232792560 {
		t.Fatal("oops")
	}
}
