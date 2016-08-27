package sol006

import "math"

func Number6(number int) int {
	sumOfSquares := 0
	squaresOfSum := 0
	for index := 1; index <= number; index++ {
		sumOfSquares += int(math.Pow(float64(index), 2))
		squaresOfSum += index
	}
	squaresOfSum = int(math.Pow(float64(squaresOfSum), 2))

	return squaresOfSum - sumOfSquares
}
