package solutions

import "testing"

// https://projecteuler.net/problem=1
func TestSolution02(t *testing.T) {
	sum := 0
	index := 1
	for {
		fib := Fibonacci(index)
		if fib <= 4000000 {
			if fib%2 == 0 {
				sum += fib
			}
			index++
			continue
		}
		break
	}
	if sum != 4613732 {
		t.Fatal("")
	}

}
