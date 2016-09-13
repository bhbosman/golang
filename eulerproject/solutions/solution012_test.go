package solutions

import (
	"fmt"
	"math"
	"testing"
)

//
// https://projecteuler.net/problem=12
//

func TestSolution12_01(t *testing.T) {
	// primes, bitarray := bhb_math.AtkinsSievePrime(104743)
	index := 2
	divisors := 0

	triangleNumber := uint64(0)
	for divisors < 500 {
		divisors = 0
		triangleNumber = uint64(index * (index + 1) / 2)
		sqrtTriangleNumber := uint64((math.Sqrt(float64(triangleNumber))))

		for i := uint64(1); i <= sqrtTriangleNumber; i++ {
			if triangleNumber%i == 0 {
				if triangleNumber/i == i {
					divisors++
				} else {
					divisors += 2
				}
			}
		}
		index++
	}
	fmt.Println(triangleNumber, divisors)
}
func TestSolution1202(t *testing.T) {
	number := FindTriangleWithXDivisiors(500)
	if number != 76576500 {
		t.Fatal("ddd")
	}
}
