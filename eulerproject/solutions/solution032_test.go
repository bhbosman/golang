package solutions

import (
	"fmt"
	"testing"

	bhb_math "github.com/bhbosman/golang/math"
)

//
// https://projecteuler.net/problem=32
//

func TestSolution32_01(t *testing.T) {

	globalData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	seqNumber := func(M int) []int {
		localData := make([]int, len(globalData))
		copy(localData, globalData)
		N := len(globalData)
		sol := make([]int, 0, len(globalData))
		for k := 1; k <= N; k++ {
			idx := (M - 1) / fact(N-k)
			sol = append(sol, localData[idx])
			// fmt.Println(sol)
			localData = RemoveIndex(localData, idx)
			M = M - idx*fact(N-k)
		}
		return sol
	}
	toInt := func(array []int) int {
		result := 0
		for i, value := range array {
			result += value * bhb_math.Pow(10, len(array)-i-1)
		}
		return result
	}
	count := 0
	sum := 0
	list := make(map[int]int)
	for i := 1; i <= fact(len(globalData)); i++ {
		iteration := seqNumber(i)
		for x := 1; x <= len(globalData)-2; x++ {
			for y := x; y+x <= len(globalData)-1; y++ {
				z := len(globalData) - x - y
				if x+y+z == len(globalData) && x <= y && y <= z {
					arrayX := iteration[0:x]
					arrayY := iteration[x : x+y]
					arrayZ := iteration[x+y : x+y+z]
					if len(arrayX)*len(arrayY) >= len(arrayZ) {
						valX := toInt(arrayX)
						valY := toInt(arrayY)
						valZ := toInt(arrayZ)
						if valX*valY == valZ {
							if _, ok := list[valZ]; !ok {
								count++
								list[valZ] = 0
								sum += valZ
								fmt.Println(count, x, y, z, iteration, arrayX, arrayY, arrayZ, valX, valY, valZ)
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(sum)
}
