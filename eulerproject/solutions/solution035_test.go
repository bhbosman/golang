package solutions

import (
	"fmt"
	"math"
	"testing"

	bhb_math "github.com/bhbosman/golang/math"
)

//
// https://projecteuler.net/problem=35
//

func TestSolution35_01(t *testing.T) {
	primeList := bhb_math.AtkinsSievePrime(1000000)
	primeListMap := make(map[int]bool)

	for _, v := range primeList {
		primeListMap[int(v)] = true
	}

	digitFactSum := func(n int) int {
		digits := math.Log(float64(n)) / math.Log(float64(10))
		return int(math.Ceil(digits))
	}

	circPrime := func(number int) bool {
		digits := digitFactSum(number)
		for i := digits - 1; i >= 0; i-- {
			if _, ok := primeListMap[number]; !ok {
				return false
			}
			value := number / bhb_math.Pow(10, digits-1)
			number = (number-(value*bhb_math.Pow(10, digits-1)))*10 + value
		}
		return true
	}

	count := 0
	for _, value := range primeList {
		if circPrime(int(value)) {
			count++
			fmt.Println(value)
		}
	}
	fmt.Println(count)
}
