package solutions

import (
	"fmt"
	"math"
	"testing"
)

import bhb_math "github.com/bhbosman/code/golang/libs/math"

//
// https://projecteuler.net/problem=34
//

func TestSolution34_01(t *testing.T) {
	maxValueforNDigits := func(n int) int {
		return n * fact(9)
	}
	digitFactSum := func(n int) int {
		digits := math.Log(float64(n)) / math.Log(float64(10))
		return int(math.Ceil(digits))
	}
	sumDigits := func(n int) int {
		leftOver := n
		result := 0
		digits := digitFactSum(n)
		for i := digits - 1; i >= 0; i-- {
			value := leftOver / bhb_math.Pow(10, i)
			result += fact(value)
			leftOver -= value * bhb_math.Pow(10, i)
		}
		return result
	}

	sum := 0

	count := 0
	for i := 2; true; i++ {
		v1 := bhb_math.Pow(10, i) - 1
		v2 := maxValueforNDigits(i)
		maxK := maxValueforNDigits(i)
		for k := v1; k <= maxK && k > bhb_math.Pow(10, i-1); k-- {
			count++
			d := sumDigits(k)
			if d == k {
				sum += k
				fmt.Println(k)
			}
		}
		if !(v1 <= v2) {
			break
		}
	}
	fmt.Println(count, sum)
}
