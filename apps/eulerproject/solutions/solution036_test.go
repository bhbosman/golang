package solutions

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"testing"

	bhb_math "github.com/bhbosman/code/golang/libs/math"
)

//
// https://projecteuler.net/problem=36
//

func TestSolution36_01(t *testing.T) {
	digitLength := func(base, n int) int {
		digits := math.Log(float64(n)) / math.Log(float64(base))
		//
		return int(math.Ceil(digits))
	}

	digrev := func(base, number int) int {
		result := 0
		nextDigit := 0
		digits := digitLength(base, number)
		startBuild := false
		for i := digits + 1; i >= 0; i-- {
			digitValueInput := bhb_math.Pow(base, i)
			value := number / digitValueInput
			if value > 0 {
				startBuild = true
			}
			if startBuild {
				number -= value * digitValueInput
				digitValueOutput := bhb_math.Pow(10, nextDigit)
				result += digitValueOutput * value
				nextDigit++
			}
		}
		//
		return result
	}

	palindome := func(digits int) []int {
		base := 10
		result := make([]int, 0, bhb_math.Pow(10, digits))
		for d := 1; d <= digits; d++ {
			if d%2 == 0 {
				m := d / 2
				startNumber := bhb_math.Pow(base, m-1)
				endNumber := bhb_math.Pow(base, m) - 1
				for n := startNumber; n <= endNumber; n++ {
					result = append(result, n*bhb_math.Pow(base, m)+digrev(base, n))
				}
			} else {
				m := (d - 1) / 2
				startNumber := bhb_math.Pow(base, m-1)
				endNumber := bhb_math.Pow(base, m) - 1
				if endNumber == 0 {
					startNumber = 0
				}
				for n := startNumber; n <= endNumber; n++ {
					for y := 0; y <= 9; y++ {
						result = append(result, n*bhb_math.Pow(base, m+1)+y*bhb_math.Pow(base, m)+digrev(base, n))
					}
				}
			}
		}
		//
		return result
	}

	base10ToBase2 := func(number int) string {
		result := ""
		base := 2
		digits := digitLength(base, number)
		startBuild := false
		for i := digits + 1; i >= 0; i-- {
			digitValueInput := bhb_math.Pow(base, i)
			value := number / digitValueInput
			if value > 0 {
				startBuild = true
			}
			if startBuild {
				number -= value * digitValueInput
				result = result + strconv.Itoa(value)
			}
		}
		//
		return result
	}

	isPalindrome := func(number string) bool {
		a := 0
		z := len(number) - 1
		for a <= z {
			if number[a] != number[z] {
				return false
			}
			a++
			z--
		}
		//
		return true
	}

	data := palindome(6)
	sort.Ints(data)
	sum := 0
	for i, k := range data {
		base2 := base10ToBase2(k)
		is := isPalindrome(base2)
		if is {
			fmt.Println(i+1, k, base2, is)
			sum += k
		}
	}
	fmt.Println(sum)
}
