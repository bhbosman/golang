package solutions

import (
	"fmt"
	"testing"

	bhb_math "github.com/bhbosman/code/golang/libs/math"
)

//
// https://projecteuler.net/problem=27
//

func TestSolution27_01(t *testing.T) {
	primes := bhb_math.AtkinsSievePrime(uint64(100000))
	isPrimeFunctionCreation := func() func(uint64) bool {
		primeMap := make(map[uint64]bool)
		for _, v := range primes {
			primeMap[v] = true
		}
		return func(value uint64) bool {
			_, ok := primeMap[value]
			return ok
		}
	}
	isPrime := isPrimeFunctionCreation()
	seq := func(a int, b uint64) (int, []int64) {
		quadraticFunction := func(n, a int, b uint64) int64 {
			return int64(n*n) - int64(a*n) + int64(b)
		}
		// derivative := func(n, a int) int64 {
		// return int64(2*n + a)
		// }

		n0 := quadraticFunction(0, 1, 41)
		// n01 := derivative(0, 1)
		if isPrime(uint64(n0)) {
			data := make([]int64, 0, 100)
			n := 0
			for {
				value := quadraticFunction(n, a, b)
				if !(isPrime(uint64(value))) {
					break
				}
				data = append(data, value)
				n++
			}
			return n, data
		}
		return 0, nil
	}
	// print := func(a int, b uint64) {
	// 	count, data := seq(a, b)
	// 	if count > 0 {
	// 		fmt.Println(a, b, count, data)
	// 	}
	// }

	// print(1, 41)
	// print(-1, 41)
	// print(-79, 1601)

	indexB := 0

	highest := -1

	for {
		b := primes[indexB]
		if b > 1000 {
			break
		}
		for a := -999; a <= 999; a++ {
			count, data := seq(a, b)
			if count > highest {
				highest = count
				if count > 1 {
					fmt.Println(a, b, count, data)
				}
			}
		}
		indexB++

	}
}
