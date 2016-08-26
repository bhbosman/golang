package solutions

// Multiplesof3and5 ...
func Multiplesof3and5(number int) int {
	result := 0
	for i := 1; i < number; i++ {

		b := false
		b = b || (i%3 == 0)
		b = b || (i%5 == 0)
		if b {
			result += i
		}
	}
	return result
}
