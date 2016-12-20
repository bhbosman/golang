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

func TestSolution37_01(t *testing.T) {
	primeList := bhb_math.AtkinsSievePrime(1000000)
	primeListMap := make(map[int]bool)
	for _, v := range primeList {
		primeListMap[int(v)] = true
	}

	digitFactSum := func(n int) int {
		digits := math.Log(float64(n)) / math.Log(float64(10))
		return int(math.Ceil(digits))
	}

	truncPrimeLeft := func(number int) ([]int, bool) {
		primes := make(map[int]bool)
		primes[number] = true

		digits := digitFactSum(number)
		for i := digits - 1; i >= 0; i-- {
			if _, ok := primeListMap[number]; !ok {
				return nil, false
			}
			pow := bhb_math.Pow(10, i)
			value := number / pow
			number -= value * pow
			if number != 0 {
				primes[number] = true
			}
		}
		result := make([]int, 0, len(primes))
		for key := range primes {
			result = append(result, key)
		}
		return result, true
	}

	truncPrimeRight := func(number int) ([]int, bool) {
		primes := make(map[int]bool)
		primes[number] = true

		digits := digitFactSum(number)
		for i := digits - 1; i >= 0; i-- {
			if _, ok := primeListMap[number]; !ok {
				return nil, false
			}

			number /= 10
			if number != 0 {
				primes[number] = true
			}
		}
		result := make([]int, 0, len(primes))
		for key := range primes {
			result = append(result, key)
		}
		return result, true
	}

	canContinue := func(number int) bool {
		switch number {
		case 2:
			return false
		case 3:
			return false
		case 5:
			return false
		case 7:
			return false
		}
		return true
	}

	count := 0
	sum := 0
	for _, value := range primeList {
		if canContinue(int(value)) {
			if primesLeft, ok := truncPrimeLeft(int(value)); ok {
				if primesRight, ok := truncPrimeRight(int(value)); ok {
					count++
					sum += int(value)
					fmt.Println(value, primesLeft, primesRight)
				}
			}
		}
	}
	fmt.Println(count, sum)
}
