package solutions

import (
	"fmt"
	"testing"

	bhb_math "github.com/bhbosman/golang/math"
)

//
// https://projecteuler.net/problem=23
//

type NumberDescription int

const (
	PerfectNumber NumberDescription = iota
	AbundantNumber
	DeficientNumber
)

func (n NumberDescription) String() string {
	switch n {
	case PerfectNumber:
		return "PerfectNumber"
	case AbundantNumber:
		return "AbundantNumber"
	case DeficientNumber:
		return "DeficientNumber"
	default:
		return "(Not Defined)"
	}
}

func TestSolution23_01(t *testing.T) {
	calculateNumberDescriptionClosure := func(maxprimes int) func(index int) (sum int, what NumberDescription) {
		primes := bhb_math.AtkinsSievePrime(uint64(maxprimes))
		return func(index int) (sum int, what NumberDescription) {
			sumDivisorsWithExclusion := func(divisors []int) int {
				result := 0
				for _, v := range divisors {
					if v != index {
						result += v
					}
				}
				return result
			}
			typeOfNumber := func(sum int) NumberDescription {
				if sum > index {
					return AbundantNumber
				} else if sum < index {
					return DeficientNumber
				}
				return PerfectNumber
			}
			primeDivisors := bhb_math.FindPrimeDivisors(primes, index)
			allPrimeDivisors := bhb_math.FindPrimeDivisorsMakeFlat(primeDivisors)
			allDivisors := bhb_math.FindAllDivisorsAndSum(allPrimeDivisors)
			sum = sumDivisorsWithExclusion(allDivisors)
			what = typeOfNumber(sum)
			//
			return sum, what
		}
	}
	calculateNumberDescription := calculateNumberDescriptionClosure(100000)
	for index := 1; index <= 50; index++ {
		sum, what := calculateNumberDescription(index)
		fmt.Println(index, sum, what)
	}
}
