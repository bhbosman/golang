package math

// Pow ...
func Pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

// Permutation ...
func Permutation(n, k int) int {
	result := 1
	for i := 0; i < k; i++ {
		result *= (n - i)
	}
	return result
}
