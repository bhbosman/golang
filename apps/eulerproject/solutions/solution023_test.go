package solutions

import (
	"testing"

	"github.com/Workiva/go-datastructures/bitarray"
	bhb_math "github.com/bhbosman/code/golang/libs/math"
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
	number := 28124
	calculateNumberDescription := calculateNumberDescriptionClosure(number)
	abundantNumbers := make([]int, 0, 10000)
	for i := 1; i <= number; i++ {
		_, whatI := calculateNumberDescription(i)
		if whatI == AbundantNumber {
			abundantNumbers = append(abundantNumbers, i)
		}
	}
	numbersNotTouched := bitarray.NewBitArray(uint64(number), true)
	for i := uint64(number); i < numbersNotTouched.Capacity(); i++ {
		numbersNotTouched.ClearBit(i)
	}
	for i, v1 := range abundantNumbers {
		for k := i; k < len(abundantNumbers); k++ {
			v2 := abundantNumbers[k]
			sum := v1 + v2
			if sum < number {
				numbersNotTouched.ClearBit(uint64(sum))
			}
			if sum >= number {
				break
			}
		}
	}
	sum := uint64(0)
	values := numbersNotTouched.ToNums()
	for _, v := range values {
		sum += v
	}
	if sum != 4179871 {
		t.Fail()
	}
}
