package solutions

import (
	"fmt"
	"testing"

	bhb_math "github.com/bhbosman/golang/math"
)

func TestSolution22_01(t *testing.T) {
	d := make(map[int]int)
	number := 10000
	primes := bhb_math.AtkinsSievePrime(uint64(number))
	for index := 1; index <= number; index++ {
		data := bhb_math.FindPrimeDivisors(primes, index)
		flat := bhb_math.FindPrimeDivisorsMakeFlat(data)
		_, sum := bhb_math.FindAllDivisorsAndSum(index, flat)
		// fmt.Println(index, data, flat, divisors, sum)
		d[index] = sum
	}
	result := make(map[int]int)
	addToResult := func(number int) {
		_, ok := result[number]
		if !ok {
			result[number] = 0
		}
	}

	for k, v := range d {
		_, ok := d[v]
		if ok {
			if k == d[v] && k != v {
				addToResult(k)
				addToResult(v)
			}
		}
	}
	sum := 0
	for k := range result {
		sum += k

	}
	fmt.Println(result, sum, sum/2)
}
