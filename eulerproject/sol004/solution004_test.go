package sol004

import "testing"

func TestSolution04(t *testing.T) {

	d := FindPalindrome()

	if d != 906609 {
		t.Fatal("oops")
	}
}
