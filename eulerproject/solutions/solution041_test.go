package solutions

import (
	"fmt"
	"testing"

	"strconv"

	bhbmath "github.com/bhbosman/golang/math"
)

//
// https://projecteuler.net/problem=40
//

func TestSolution41_01(t *testing.T) {
	primes := bhbmath.AtkinsSievePrime(uint64(bhbmath.Pow(10, 8)))

	isPandigital := func(s string) bool {
		count := len(s)
		data := make([]bool, count, count)
		for _, v := range s {
			idx := int(v - '0' - 1)
			if idx >= 0 && idx < len(data) {
				if data[idx] {
					return false
				}
				data[idx] = true
				count--
			}
		}
		return count == 0
	}
	for _, p := range primes {
		s := strconv.Itoa(int(p))
		if isPandigital(s) {
			fmt.Println(s)
		}
	}
}
