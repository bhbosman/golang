package solutions

import (
	"fmt"
	"math"
	"testing"
)

import bhb_math "github.com/bhbosman/golang/math"

//
// https://projecteuler.net/problem=29s
//

type AB struct {
	a, b int
}

func TestSolution29_01(t *testing.T) {
	primes := bhb_math.AtkinsSievePrime(uint64(100000))

	run3 := func(minValue, maxValue int) int {
		// result := make([]int, 0, Pow(maxValue-minValue+1, 2))
		result := 0
		mainExclusionList := make(map[AB]int)
		between := func(value int) bool {
			return minValue <= value && value <= maxValue
		}

		dealWithIteration := func(excl map[AB]int, a, b int) {
			primedataForBValue := bhb_math.FindPrimeDivisors(primes, b)
			primesFlatForBValue := bhb_math.FindPrimeDivisorsMakeFlat(primedataForBValue)
			divsB := bhb_math.FindAllDivisorsAndSum(primesFlatForBValue)
			// fmt.Println(a, b, divsB)
			for _, valueB := range divsB {
				i, k := bhb_math.Pow(a, valueB), b/valueB
				if between(i) && between(k) {
					mapdata := AB{a: i, b: k}
					if _, ok := excl[mapdata]; !ok {
						excl[mapdata] = 0
					}
				}
			}
		}
		for i := minValue; i <= maxValue; i++ {
			for k := minValue; k <= maxValue; k++ {
				iterPoint := AB{a: i, b: k}
				if _, ok := mainExclusionList[iterPoint]; !ok {
					//Pow(i, p) <= Pow(maxValue, maxValue)

					for p := 1; math.Pow(float64(i), float64(p)/float64(maxValue)) <= float64(maxValue); p++ {
						// value := Pow(Pow(i, p), k)
						iterExclusionList := make(map[AB]int)
						dealWithIteration(iterExclusionList, i, k*p)
						added := false
						for key := range iterExclusionList {
							if _, ok := mainExclusionList[key]; !ok {
								mainExclusionList[key] = 0
								added = true
							}
						}
						if added {
							result++
						}
					}
				}
			}
		}
		// sort.Ints(result)
		return result
	}

	valueA := 2
	valueB := 100

	// fmt.Println(run1(valueA, valueB))

	print := func(values int) {
		fmt.Println(values)
	}

	print(run3(valueA, valueB))
	// print(bruteForce(valueA, valueB))
}
func TestSolution29_02(t *testing.T) {
	// minValue := 2
	// maxValue := 100
	max := 100

	init := func() [][]bool {
		duplicates := make([][]bool, max, max)
		for index := range duplicates {
			entry := make([]bool, max, max)
			duplicates[index] = entry
		}
		return duplicates
	}

	duplicates := init()

	mark_duplicates := func(a1 int, b1 int) {
		for b2 := 2; b2 <= max; b2++ {
			b3 := b1 * b2
			for f1 := 1; f1 < b1; f1++ {
				if b3%f1 == 0 {
					f2 := b3 / f1
					a2 := bhb_math.Pow(a1, f1)
					if a2 <= max {
						if f2 <= max {
							duplicates[a2-1][f2-1] = true
						}
					}
				}
			}
		}
	}

	sweep := func(a int) {
		for b := 2; b <= max; b++ {
			c := bhb_math.Pow(a, b)
			if c > max {
				break
			}
			mark_duplicates(a, b)
		}
	}

	count_nondups := func() int {
		count := 0
		for a := 2; a <= max; a++ {
			for b := 2; b <= max; b++ {
				if !duplicates[a-1][b-1] {
					count++
				}
			}
		}

		return count
	}

	//fmt.Println(duplicates)

	for a := 2; a <= max; a++ {
		sweep(a)
	}

	fmt.Println(count_nondups())

}

func TestSolution29_03(t *testing.T) {
	minValue := 2
	maxValue := 100
	combinations := make(map[AB]int)
	init := func() {
		for a := minValue; a <= maxValue; a++ {
			for b := minValue; b <= maxValue; b++ {
				combinations[AB{a: a, b: b}] = 0
			}
		}
	}
	run := func() {
		findDuplicates := func(inputA, inputB int) {
			for iter := minValue; iter <= maxValue; iter++ {
				number := inputB * iter
				for d := 1; d < inputB; d++ {
					if number%d == 0 {
						a := number / d
						b := bhb_math.Pow(inputA, d)
						delete(combinations, AB{a: a, b: b})
					}
				}
			}
		}
		for a := minValue; a <= maxValue; a++ {
			for b := minValue; b <= maxValue; b++ {
				c := bhb_math.Pow(a, b)
				if c > maxValue {
					break
				}
				findDuplicates(a, b)
			}
		}
	}
	init()
	run()
	fmt.Println(combinations, len(combinations))
}
