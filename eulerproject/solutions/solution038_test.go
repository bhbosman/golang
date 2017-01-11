package solutions

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

// "fmt"

// "sort"
// "testing"

// bhb_math "github.com/bhbosman/golang/math"

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
	min := func(value float64) int {
		if value == 0 {
			return 1
		}
		return int(value)
	}
	index := 1
	for {
		multiplier := 1
		numbers := 0
		s := ""
		for {
			value := index * multiplier
			s = s + strconv.Itoa(value)
			digits := min(math.Ceil(math.Log10(float64(value))))
			numbers += int(digits)
			if numbers == 9 {
				if isPandigital(s) {
					fmt.Println(index, multiplier, value, digits, numbers, s)
				}
			} else if numbers > 9 {
				break
			}
			multiplier++
			if multiplier > 9 {
				break
			}
		}
		index++
		if index > 10000 {
			break
		}
	}
}
