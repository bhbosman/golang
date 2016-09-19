package solutions

import (
	"fmt"
	"testing"

	bhb_math "github.com/bhbosman/golang/math"
)

//
// https://projecteuler.net/problem=21
//

func FindPrimeDivisors(primes []uint64, number int) map[int]int {
	result := make(map[int]int)
	leftOver := number
	for _, prime := range primes {
		count := 0
		for leftOver%int(prime) == 0 {
			count++
			leftOver /= int(prime)
		}
		if count > 0 {
			result[int(prime)] = count
		}
		if leftOver == 1 {
			break
		}
	}
	return result
}

func FindPrimeDivisorsMakeFlat(data map[int]int) []int {
	result := make([]int, 0, 100)
	for k, v := range data {
		for index := uint64(1); index <= uint64(v); index++ {
			result = append(result, k)
		}
	}
	return result
}
func findAllDivisorsAndSum(number int, data []int) ([]int, int) {
	divisors := make(map[int]int)
	divisors[1] = 0

	checkAndAdd := func(value int, temp map[int]int) {
		if value != number {
			_, ok := divisors[int(value)]
			if !ok {
				temp[value] = 0
			}
		}
	}
	for _, v := range data {
		temp := make(map[int]int)

		for k := range divisors {
			checkAndAdd(k*v, temp)
		}
		checkAndAdd(v, temp)
		for k := range temp {
			divisors[k] = 0
		}
	}
	result := make([]int, 0, len(divisors))
	sum := 0
	for k := range divisors {
		result = append(result, k)
		sum += k
	}
	return result, sum
}

func TestSolution21_01(t *testing.T) {
	d := make(map[int]int)
	number := 10000
	primes := bhb_math.AtkinsSievePrime(uint64(number))
	for index := 1; index <= number; index++ {
		data := FindPrimeDivisors(primes, index)
		flat := FindPrimeDivisorsMakeFlat(data)
		_, sum := findAllDivisorsAndSum(index, flat)
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
