package math

// FindPrimeDivisors ...
func FindPrimeDivisors(primes []uint64, number int) map[int]int {
	result := make(map[int]int)
	leftOver := number
	for _, prime := range primes {
		count := 0
		for leftOver%int(prime) == 0 {
			count++
			leftOver /= int(prime)
		}
		if count > 0 {
			result[int(prime)] = count
		}
		if leftOver == 1 {
			break
		}
	}
	return result
}
