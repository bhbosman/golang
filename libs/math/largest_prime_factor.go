package math

import "math"

func LargestPrimeFactor(number uint64) uint64 {
	lastFactor := uint64(1)
	if number%2 == 0 {
		lastFactor = 2
		number = number / 2
		for number%2 == 0 {
			number = number / 2
		}
	}

	factor := uint64(3)
	maxFactor := uint64(math.Sqrt(float64(number)))

	for (number > 1) && (factor < maxFactor) {
		if number%factor == 0 {
			number = number / factor
			for number%factor == 0 {
				number = number / factor
			}
			maxFactor = uint64(math.Sqrt(float64(number)))
		}
		factor = factor + 2
	}
	if number == 1 {
		return lastFactor
	}
	return number
}
