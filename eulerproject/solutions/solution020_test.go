package solutions

import (
	"math/big"
	"testing"
)

func TestSolution20_01(t *testing.T) {

	FactorialClosure := func() func(n int) *big.Int {
		a, b := big.NewInt(1), big.NewInt(1)
		return func(n int) *big.Int {
			if n > 1 {
				a, b = b, b.Mul(b, big.NewInt(int64(n)))
			} else {
				return big.NewInt(1)
			}

			return b
		}
	}
	dd := FactorialClosure()

	s := ""
	for i := 1; i <= 100; i++ {
		n := dd(i)
		s = n.String()
	}
	sum := 0
	for _, r := range s {
		sum += int(r - '0')
	}
	if 648 != sum {
		t.Fail()
	}
}
