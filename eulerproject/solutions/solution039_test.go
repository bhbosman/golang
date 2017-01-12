package solutions

import (
	"fmt"
	"testing"
)

//
// https://projecteuler.net/problem=39
//

func TestSolution39_01(t *testing.T) {
	perimeter := func(p int) int {
		result := 0
		for x := 1; x <= p-2; x++ {
			for y := x; x+y <= p-1; y++ {
				z := p - x - y
				if x*x+y*y == z*z {
					result++
					// fmt.Println(x, y, z, x+y+z)
				}
			}
		}
		return result
	}
	high := 0
	for index := 0; index <= 1000; index++ {
		c := perimeter(index)
		if c > high {
			high = c
			fmt.Println(index, c)
		}
	}
}
