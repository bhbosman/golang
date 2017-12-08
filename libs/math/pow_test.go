package math

import (
	"testing"
)

func TestPermutation(t *testing.T) {
	testPermutation := func(n, k, ans int) {
		a := Permutation(n, k)
		if a != ans {
			t.Fatalf("%dP%d failed. Actual: %d, Wanted: %d", n, k, a, ans)
		}
	}

	testPermutation(15, 4, 32760)
	testPermutation(9, 2, 72)
	testPermutation(9, 0, 1)
}
