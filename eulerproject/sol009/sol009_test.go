package sol009

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"testing"
)

//
// https://projecteuler.net/problem=9
//

//
// My solution
//
func TestSolution09_01(t *testing.T) {
	sumOfLength := 1000
	minLenghtOfA := int((sumOfLength - 3) / 3)
	finished := false

	beginValue := 1
	for aa, a, jumpaa := beginValue*beginValue, beginValue, 2*beginValue+1; !finished && a <= minLenghtOfA; aa, a, jumpaa = aa+jumpaa, a+1, jumpaa+2 {
		for bb, b, jumpbb := aa+jumpaa, a+1, 2*(a+1)+1; !finished && a+b+b+1 <= sumOfLength; bb, b, jumpbb = bb+jumpbb, b+1, jumpbb+2 {
			c := sumOfLength - b - a
			cc := c * c
			if aa+bb == cc {
				// finished = true
				fmt.Println(a, b, c, a*b*c)
			}
		}
	}
}

//
// Given solution
//
func gcd(a, b uint64) *big.Int {
	buffer := make([]byte, 8, 8)

	binary.LittleEndian.PutUint64(buffer, uint64(a))
	bigIntA := &big.Int{}
	bigIntA.SetBytes(buffer)

	binary.LittleEndian.PutUint64(buffer, uint64(b))
	bigIntB := &big.Int{}
	bigIntB.SetBytes(buffer)

	gcd := bigIntA.GCD(nil, nil, bigIntA, bigIntB)
	return gcd
}

func compare(gcd *big.Int, b uint64) bool {

	buffer := make([]byte, 8, 8)
	binary.LittleEndian.PutUint64(buffer, uint64(b))
	bigIntB := &big.Int{}
	bigIntB.SetBytes(buffer)

	return gcd.Cmp(bigIntB) == 0
}

func TestSolution09_02(t *testing.T) {
	s := 1000

	s2 := s / 2
	mlimit := math.Ceil(math.Sqrt(float64(s)))

	for m := 2; m <= int(mlimit); m++ {
		if s2%m == 0 {
			sm := s2 / m
			for sm%2 == 0 { // reduce the search space by
				sm = sm / 2 // removing all factors 2
			}
			k := 0
			if m%2 == 1 {
				k = m + 2
			} else {
				k = m + 1
			}
			for k < 2*m && k <= sm {
				if sm%k == 0 && compare(gcd(uint64(k), uint64(m)), 1) {
					d := s2 / (k * m)
					n := k - m
					a := d * (m*m - n*n)
					b := 2 * d * m * n
					c := d * (m*m + n*n)
					fmt.Println(a, b, c, m, n, d)
				}
				k = k + 2
			}
		}
	}

}
