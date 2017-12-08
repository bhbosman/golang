package solutions

import (
	"fmt"
	"math/big"
	"testing"
)

//
// https://projecteuler.net/problem=23
//

func createFib() func() (int, *big.Int) {
	index := 0
	a := big.NewInt(int64(0))
	b := big.NewInt(int64(1))
	return func() (int, *big.Int) {
		index++
		a, b = b, a.Add(a, b)
		return index, a
	}
}

func TestSolution25_01(t *testing.T) {
	fib := createFib()
	index := 0
	for {
		index++
		// if index > 10000 {
		// 	break
		// }
		a, b := fib()
		s := b.String()

		if len(s) >= 1000 {
			fmt.Println(a, b, s, len(s))
			break
		}
	}
}
