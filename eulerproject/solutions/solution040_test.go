package solutions

import (
	"fmt"
	"testing"

	"strconv"

	bhbmath "github.com/bhbosman/golang/math"
)

//
// https://projecteuler.net/problem=40
//

func TestSolution40_01(t *testing.T) {
	ChampernowneConstant := func(index int) byte {
		leftOver := index
		power := 0
		for {
			minPowerValue := bhbmath.Pow(10, power)
			maxPowerValue := bhbmath.Pow(10, power+1) - 1
			valuesCovered := maxPowerValue - minPowerValue + 1
			stepLength := power + 1
			indexesCovered := valuesCovered * (stepLength)
			if leftOver <= indexesCovered {
				a := (leftOver - 1) / stepLength
				steps := leftOver / stepLength
				valueToUse := minPowerValue + a
				i := (power + leftOver - steps*stepLength) % stepLength
				// fmt.Println(index, power, indexesCovered, leftOver, a, steps, valueToUse, i)
				s := strconv.Itoa(valueToUse)
				return s[i] - '0'
			}
			power++
			leftOver -= indexesCovered
		}
	}
	ans := ChampernowneConstant(1) * ChampernowneConstant(10) * ChampernowneConstant(100) * ChampernowneConstant(1000) * ChampernowneConstant(10000) * ChampernowneConstant(100000) * ChampernowneConstant(1000000)
	fmt.Println(ans)
}
