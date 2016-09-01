package sol008

import (
	"fmt"
	"math"
	"testing"
)

//
// https://projecteuler.net/problem=8
//

var strData = ("73167176531330624919225119674426574742355349194934" +
	"96983520312774506326239578318016984801869478851843" +
	"85861560789112949495459501737958331952853208805511" +
	"12540698747158523863050715693290963295227443043557" +
	"66896648950445244523161731856403098711121722383113" +
	"62229893423380308135336276614282806444486645238749" +
	"30358907296290491560440772390713810515859307960866" +
	"70172427121883998797908792274921901699720888093776" +
	"65727333001053367881220235421809751254540594752243" +
	"52584907711670556013604839586446706324415722155397" +
	"53697817977846174064955149290862569321978468622482" +
	"83972241375657056057490261407972968652414535100474" +
	"82166370484403199890008895243450658541227588666881" +
	"16427171479924442928230863465674813919123162824586" +
	"17866458359124566529476545682848912883142607690042" +
	"24219022671055626321111109370544217506941658960408" +
	"07198403850962455444362981230987879927244284909188" +
	"84580156166097919133875499200524063689912560717606" +
	"05886116467109405077541002256983155200055935729725" +
	"71636269561882670428252483600823257530420752963450")

func getData() []byte {

	n := 1
	intData := make([]byte, 0, n*1000)
	for i := 1; i <= n; i++ {
		for _, r := range strData {

			intData = append(intData, byte(r-'0'))
		}
	}
	fmt.Println(len(intData))

	return intData
}

func TestSolution06_01(t *testing.T) {
	largest := int64(math.MinInt64)
	intData := getData()

	lastZero := 0
	backTrack := 13
	number := int64(1)
	for i, v := range intData {
		// prevNumber := number
		remove := int64(0)
		if i-backTrack >= 0 {
			remove = int64(intData[i-backTrack])
		} else {
			remove = 1
		}
		rans := int64(0)

		removeCan := false
		if i-lastZero-backTrack >= 0 {
			removeCan = true
		}

		if remove != 0 {
			if removeCan {
				rans = number / remove
				number = rans
			} else {
				rans = -1
			}
		} else {
			rans = -1
		}
		if v != 0 {
			number *= int64(v)
		} else {
			number = 1
			lastZero = i + 1
		}
		if number > largest {
			largest = number
		}

		// fmt.Println("Index:", i, intData[int(math.Max(float64(i-backTrack+1), float64(0))):i+1], prevNumber, remove, rans, v, number, lastZero, removeCan, largest)
	}
	fmt.Println(largest)
}

func TestSolution06_02(t *testing.T) {
	backTrack := 13
	largest := int64(math.MinInt64)
	acc := int64(1)
	consumed := 0

	intData := getData()

	values := make([]byte, backTrack, backTrack)

	for i, v := range intData {
		if consumed < backTrack {
			consumed++
		} else {
			acc /= int64(values[i%backTrack])
		}

		values[i%backTrack] = v
		acc *= int64(values[i%backTrack])

		if acc > largest && consumed == backTrack {
			largest = acc
		}

		if acc == 0 {
			// We read a zero; reset the consumed buffer and accumulator.
			consumed, acc = 0, int64(1)
		}

	}
	fmt.Println(largest)
}
