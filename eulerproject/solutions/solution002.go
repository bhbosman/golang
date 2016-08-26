package solutions

// Fibonacci ...
func Fibonacci(number int) int {
	switch number {
	case 1:
		return 1
	case 2:
		return 2
	default:
		return Fibonacci(number-1) + Fibonacci(number-2)
	}
}
