package solutions

import "testing"
import "math"

//
// https://projecteuler.net/problem=6
//
func Number6MyOne(number int) int {
	sumOfSquares := 0
	squaresOfSum := 0
	for index := 1; index <= number; index++ {
		sumOfSquares += int(math.Pow(float64(index), 2))
		squaresOfSum += index
	}
	squaresOfSum = int(math.Pow(float64(squaresOfSum), 2))

	return squaresOfSum - sumOfSquares
}

func Number06Answer(number int) int {
	limit := 100
	sum := limit * (limit + 1) / 2
	sumsq := (2*limit + 1) * (limit + 1) * limit / 6
	return sum*sum - sumsq
}

const number = 100

func TestSolution06_01(t *testing.T) {

	n := Number6MyOne(number)

	if n != 25164150 {
		t.Fatal("oops")
	}
}
func TestSolution06_02(t *testing.T) {

	n := Number06Answer(number)

	if n != 25164150 {
		t.Fatal("oops")
	}
}
