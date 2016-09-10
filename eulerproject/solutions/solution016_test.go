package solutions

import (
	"fmt"
	"math/big"
	"testing"
)

//
// https://projecteuler.net/problem=15
//

func TestSolution16_01(t *testing.T) {
	ans := big.NewInt(2)

	ans = ans.Lsh(ans, 999)
	s := ans.String()
	fmt.Println(s)
	n := 0
	for _, r := range s {
		v := int(r - '0')
		n += v
		fmt.Println(r, v, n)
	}

	fmt.Println(n)

}
