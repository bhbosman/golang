package solutions

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

//
// https://projecteuler.net/problem=38
//

func TestSolution38_01(t *testing.T) {
	isPandigital := func(s string) bool {
		count := 9
		data := make([]bool, 9, 9)
		for _, v := range s {
			idx := v - '0' - 1
			if idx >= 0 {
				if data[idx] {
					return false
				}
				data[idx] = true
				count--
			}
		}
		return count == 0
	}
	minStringCount := func(value int) int {
		if value == 0.0 {
			return 1
		}
		return value
	}
	index := 1
	for {
		numbers := 0
		s := ""
		for multiplier := 1; multiplier <= 9; multiplier++ {
			value := index * multiplier
			s = s + strconv.Itoa(value)
			digits := minStringCount(int(math.Ceil(math.Log10(float64(value)))))
			numbers += digits
			if numbers == 9 {
				if isPandigital(s) {
					fmt.Println(index, multiplier, value, digits, numbers, s)
				}
			} else if numbers > 9 {
				break
			}
		}
		index++
		if index > 10000 {
			break
		}
	}
}
