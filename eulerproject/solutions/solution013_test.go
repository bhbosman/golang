package solutions

import (
	"math"
	"math/big"
	"strconv"
	"testing"
)

//
// https://projecteuler.net/problem=13
//

func TestSolution13_01(t *testing.T) {
	bigNumbers, longest, _ := readFileData("solution013_data.txt")
	revAnswer := make([]byte, 0, len(bigNumbers))
	overflow := 0
	for idx := longest - 1; idx >= 0; idx-- {
		values := overflow
		for _, line := range bigNumbers {
			c := int(line[idx] - '0')
			values += c
		}
		digit := values % 10
		overflow = int(math.Trunc(float64(values / 10)))
		revAnswer = append(revAnswer, byte(digit+'0'))
	}
	ans := make([]byte, len(revAnswer), len(revAnswer))
	for i, b := range revAnswer {
		ans[len(ans)-i-1] = b
	}
	TheAnswer := strconv.Itoa(overflow) + string(ans)

	if TheAnswer != "5537376230390876637302048746832985971773659831892672" {
		t.Fail()
	}
}

func TestSolution13_02(t *testing.T) {
	bigNumbers, _, _ := readFileData("solution013_data.txt")
	sum := big.NewInt(0)
	for _, line := range bigNumbers {
		number := big.NewInt(0)
		number.SetString(line, 10)
		sum = sum.Add(sum, number)
	}
	TheAnswer := sum.String()

	if TheAnswer != "5537376230390876637302048746832985971773659831892672" {
		t.Fail()
	}
}
