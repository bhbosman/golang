package sol004

import "testing"

//
// https://projecteuler.net/problem=4
//

func TestSolution04(t *testing.T) {

	d := FindPalindrome()

	if d != 906609 {
		t.Fatal("oops")
	}
}
