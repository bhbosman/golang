package sol006

import "testing"

//
// https://projecteuler.net/problem=6
//

func TestSolution06(t *testing.T) {

	n := Number6(100)

	if n != 25164150 {
		t.Fatal("oops")
	}
}
